package pg

import (
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(entities.Simulator{})
}