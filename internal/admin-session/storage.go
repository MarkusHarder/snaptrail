package adminsession

import (
	"context"
	"snaptrail/internal/db"
	"snaptrail/internal/structs"

	"gorm.io/gorm"
)

type repository interface {
	getSessions(public bool) (sessions []structs.Session, err error)
	getSessionById(id string) (session structs.Session, err error)
	createOrUpdateSession(session *structs.Session) (err error)
	createOrUpdateThumbnail(img *structs.Thumbnail) (err error)
	deleteSession(id string) (err error)
}

func newRepo(ctx context.Context) repository {
	return repo{
		db: db.GetDb().WithContext(ctx),
	}
}

func withTx(tx *gorm.DB, ctx context.Context) repository {
	return repo{db: tx.WithContext(ctx)}
}

type repo struct {
	db *gorm.DB
}

func (r repo) getSessions(public bool) (sessions []structs.Session, err error) {
	query := r.db.Model(&structs.Session{})
	if public {
		query = query.Where("published = true")
	}
	err = query.Preload("Thumbnail").Order("created_at ASC").Find(&sessions).Error
	return
}

func (r repo) getSessionById(id string) (session structs.Session, err error) {
	err = r.db.Preload("Thumbnail").Where("id = ?", id).Find(&session).Error
	return
}

func (r repo) createOrUpdateSession(session *structs.Session) (err error) {
	err = r.db.Save(&session).Error
	return
}

func (r repo) deleteSession(id string) (err error) {
	err = r.db.Delete(structs.Session{ID: id}).Error
	return
}

func (r repo) createOrUpdateThumbnail(img *structs.Thumbnail) (err error) {
	err = r.db.Save(img).Error
	return
}
