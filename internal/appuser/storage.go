package appuser

import (
	"snaptrail/internal/db"
	"snaptrail/internal/structs"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type repository interface {
	getUserByName(username string) (user structs.User, err error)
	changeUserPassword(userId uint64, newPassword string, newVersion int64) (err error)
}

func newRepo() repository {
	return repo{
		db: db.GetDb(),
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) getUserByName(username string) (user structs.User, err error) {
	err = r.db.Where("username = ?", username).First(&user).Error
	return
}

func (r repo) changeUserPassword(userId uint64, newPassword string, newVersion int64) (err error) {
	log.Info().Msgf("Got id: %d, pw: %s, version: %d", userId, newPassword, newVersion)

	return r.db.Model(&structs.User{}).Where("ID = ?", userId).Updates(structs.User{Password: newPassword, Version: newVersion}).Error
}
