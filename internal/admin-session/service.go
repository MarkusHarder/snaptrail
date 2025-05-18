package adminsession

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"snaptrail/internal/db"
	"snaptrail/internal/extractor"
	"snaptrail/internal/structs"
	"strings"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Service interface {
	GetSessions(public bool) (session []structs.Session, err error)
	createOrUpdateSession(session *structs.Session, img *multipart.FileHeader) (err error)
	deleteSession(id uint64) (err error)
}

func NewService(ctx context.Context) Service {
	return svc{
		ctx:  ctx,
		repo: newRepo(ctx),
	}
}

type svc struct {
	ctx    context.Context
	repo   repository
	txRepo repository
}

func (s svc) GetSessions(public bool) (session []structs.Session, err error) {
	return s.repo.getSessions(public)
}

func (s svc) createOrUpdateSession(session *structs.Session, img *multipart.FileHeader) (err error) {
	thumbnail, err := createThumbnail(img)
	if err != nil {
		log.Err(err).Msg("could not create thumbnail object for database from form")
		return err
	}
	tx := db.GetDb().Begin()
	txRepo := s.getTxRepo(tx)
	err = txRepo.createOrUpdateSession(session)
	if err != nil {
		tx.Rollback()
		return err
	}

	thumbnail.SessionID = session.ID
	err = txRepo.createOrUpdateThumbnail(*thumbnail)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (s svc) deleteSession(id uint64) (err error) {
	return s.repo.deleteSession(id)
}

func (s svc) getTxRepo(tx *gorm.DB) repository {
	if s.txRepo != nil {
		return s.txRepo
	}
	return withTx(tx, s.ctx)
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
