package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDb(){
	str:= "host=localhost port=5432 user=admin dbname=civoo sslmode=disable password=root"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	
	if err != nil {
		log.Fatal("error:", err)
	}
	
	db = database
	
	config, _ := db.DB()
	
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	RunMigrations(db)
}


func GetDatabase() *gorm.DB {
 return db
}