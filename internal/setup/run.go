package setup

import (
	"snaptrail/internal/config"
	"snaptrail/internal/db"
	"snaptrail/internal/structs"

	"github.com/rs/zerolog/log"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func CreateAdmin() error {
	repo := db.GetDb()

	username := config.Get().AdminUsername

	admin, err := getAdminUser(repo, username)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal().Msg("failed checking the db for admin user, quitting")
	}

	if err == gorm.ErrRecordNotFound {

		log.Info().Msgf("could not find existing user with username: %s, creating new user...", username)

		password := config.Get().AdminPassword

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal().Msg("failed generating password hash")
		}

		user := structs.User{
			Username: username,
			Password: string(hash[:]),
			Version:  0,
			Role:     structs.AdminRole,
		}

		return createUser(repo, &user)

	} else {

		log.Info().Msgf("found admin user with username: %s", admin.Username)

		return nil

	}
}

func createUser(repo *gorm.DB, user *structs.User) error {
	return repo.Create(&user).Error
}

func getAdminUser(repo *gorm.DB, username string) (user structs.User, err error) {
	err = repo.Where("username = ?", username).First(&user).Error

	return
}
