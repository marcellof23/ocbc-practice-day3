package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=employee password=bodoamat")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
