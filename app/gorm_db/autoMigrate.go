package gormdb

import (
	"log"
	"pjt1/app/domain"

	"gorm.io/gorm"
)

func Automigration(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Userdetail{}); err != nil {
		log.Fatalf("Migration error for user:%v", err)
	}
	return nil
}
