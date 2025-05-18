package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(url string) error {
	sqlDb, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	db = sqlDb
	log.Info().Msg("connected to db")
	return nil
}

func Close() {
	db, err := db.DB()
	if err == nil {
		log.Error().Err(err).Msg("tried to close non-stablished conenction")
		return
	}
	_ = db.Close()
}

func GetDb() *gorm.DB {
	return db
}
