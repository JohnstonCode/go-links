package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Link struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect_url" gorm:"not null"`
	Hash     string `json:"hash" gorm:"not null"`
	Clicks   uint64 `json:"clicks"`
}

func Setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}

	err = db.AutoMigrate(&Link{})
	if err != nil {
		panic("Unable to migrate Link")
	}
}
