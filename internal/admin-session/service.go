package adminsession

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"snaptrail/internal/config"
	"snaptrail/internal/db"
	"snaptrail/internal/extractor"
	"snaptrail/internal/s3"
	"snaptrail/internal/structs"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// make this a bit longer than jwt TTL
const validSeconds = int64(35 * time.Minute / time.Second)

type Service interface {
	GetSessions(public bool) (session []structs.Session, err error)
	createSession(session *structs.Session, img *multipart.FileHeader) (err error)
	updateSession(session *structs.Session, img *multipart.FileHeader) (err error)
	deleteSession(id string) (err error)
}

func NewService(ctx context.Context) Service {
	return svc{
		ctx:    ctx,
		repo:   newRepo(ctx),
		bucket: s3.NewBucketBasics(),
	}
}

type svc struct {
	ctx    context.Context
	repo   repository
	txRepo repository
	bucket s3.BucketBasics
}

func (s svc) GetSessions(public bool) (sessions []structs.Session, err error) {
	sessions, err = s.repo.getSessions(public)
	if err != nil {
		return sessions, err
	}

	for i, sess := range sessions {
		req, err := s.bucket.GetObject(context.Background(), config.Get().S3Bucket, sess.Thumbnail.Path, validSeconds)
		if err != nil {
			log.Err(err).Msgf("failed to get presigned url for session: %s with thumbnnail:%s", sess.Name, sess.Thumbnail.Filename)
		}

		sessions[i].Thumbnail.ImageSrc = req.URL
	}

	return
}

func (s svc) createSession(session *structs.Session, img *multipart.FileHeader) (err error) {
	thumbnail, err := createThumbnail(img)
	if err != nil {
		log.Err(err).Msg("could not create thumbnail object for database from form")
		return err
	}

	tx := db.GetDb().Begin()
	txRepo := s.getTxRepo(tx)
	err = txRepo.createOrUpdateSession(session)
	if err != nil {
		log.Err(err).Msg("failed creating or updating session in transaction")
		tx.Rollback()
		return err
	}

	thumbnail.SessionID = session.ID
	thumbnail.ID = uuid.New().String()
	thumbnail.Path = fmt.Sprintf("%s/%s", thumbnail.ID, thumbnail.Filename)
	err = txRepo.createOrUpdateThumbnail(thumbnail)
	if err != nil {
		log.Err(err).Msg("failed creating or updating thumbnail in transaction")
		tx.Rollback()
		return err
	}

	err = s.bucket.UploadFile(context.Background(), config.Get().S3Bucket, *thumbnail)
	if err != nil {
		log.Err(err).Msg("failed uploading the image to s3 in transaction")
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s svc) updateSession(session *structs.Session, img *multipart.FileHeader) (err error) {
	thumbnail, err := createThumbnail(img)
	if err != nil {
		log.Err(err).Msg("could not create thumbnail object for database from form")
		return err
	}

	oldSession, err := s.repo.getSessionById(session.ID)
	if err != nil {
		return err
	}

	oldKey := oldSession.Thumbnail.Path
	thumbnail.ID = oldSession.Thumbnail.ID
	thumbnail.SessionID = session.ID
	thumbnail.Path = fmt.Sprintf("%s/%s", thumbnail.ID, thumbnail.Filename)

	tx := db.GetDb().Begin()
	txRepo := s.getTxRepo(tx)
	err = txRepo.createOrUpdateSession(session)
	if err != nil {
		log.Err(err).Msg("failed creating or updating session in transaction")
		tx.Rollback()
		return err
	}

	err = txRepo.createOrUpdateThumbnail(thumbnail)
	if err != nil {
		log.Err(err).Msg("failed creating or updating thumbnail in transaction")
		tx.Rollback()
		return err
	}

	err = s.bucket.UploadFile(context.Background(), config.Get().S3Bucket, *thumbnail)
	if err != nil {
		log.Err(err).Msg("failed uploading the image to s3 in transaction")
		tx.Rollback()
		return err
	}

	if oldKey != thumbnail.Path {
		err = s.bucket.DeleteObjects(context.Background(), config.Get().S3Bucket, []string{oldKey})
		if err != nil {
			log.Err(err).Msg("failed deleting old image on s3 after update")
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (s svc) deleteSession(id string) (err error) {
	session, err := s.repo.getSessionById(id)
	if err != nil {
		return err
	}

	err = s.bucket.DeleteObjects(context.Background(), config.Get().S3Bucket, []string{session.Thumbnail.Path})
	if err != nil {
		return
	}

	return s.repo.deleteSession(id)
}

func createThumbnail(img *multipart.FileHeader) (t *structs.Thumbnail, err error) {
	file, err := img.Open()
	if err != nil {
		return nil, err
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	filetype := http.DetectContentType(fileBytes)
	if !strings.HasPrefix(filetype, "image") {
		return nil, fmt.Errorf("invalid filetype for thumbnail")
	}

	exifMD, err := extractor.CreateExifMetadataFromImage(fileBytes)
	if err != nil {
		return nil, err
	}

	thumbnail := structs.Thumbnail{
		Filename:     img.Filename,
		MimeType:     filetype,
		ExifMetadata: *exifMD,
		Data:         fileBytes,
	}

	return &thumbnail, nil
}

func (s svc) getTxRepo(tx *gorm.DB) repository {
	if s.txRepo != nil {
		return s.txRepo
	}
	return withTx(tx, s.ctx)
}
