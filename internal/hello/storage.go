package hello

import (
	"snaptrail/internal/db"
	"snaptrail/internal/structs"

	"gorm.io/gorm"
)

type repository interface {
	getHello() (hello structs.Hello, err error)
}

func newRepo() repository {
	return repo{
		db: db.GetDb(),
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) getHello() (hello structs.Hello, err error) {
	err = r.db.First(&hello).Error
	return
}
