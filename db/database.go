package db

import (
	"log"
	"github.com/sammatcha/urlshortener/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbURL := "postgres://yourname:password@localhost:5432/urlshortener"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to db:", err)
		panic("failed to connect to database")
	}
	log.Println("Successfully connected to db")

	err = db.AutoMigrate(&models.Url{})
	if err != nil{
		log.Println("failed to migrate url db", err)
		return nil, err
	}
	return db, err
}