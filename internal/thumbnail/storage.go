package thumbnail

import (
	"snaptrail/internal/db"
	"snaptrail/internal/structs"

	"gorm.io/gorm"
)

type repository interface {
	getThumbnailById(sessionId uint64, thumbnailId uint64, published bool) (thumbnail structs.Thumbnail, err error)
}

func newRepo() repository {
	return repo{
		db: db.GetDb(),
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) getThumbnailById(sessionId uint64, thumbnailId uint64, published bool) (thumbnail structs.Thumbnail, err error) {
	// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
	query := r.db.Where("thumbnails.id = ?", thumbnailId)
	if published {
		query = query.Joins("JOIN sessions ON sessions.id = ?", sessionId).Where("sessions.published = true")
	}
	err = query.Find(&thumbnail).Error
	return
}
