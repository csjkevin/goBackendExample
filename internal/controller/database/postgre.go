package database

import (
	"github.com/jinzhu/gorm"
	"goBackendExample/internal/model"
	"log"
)

func GetPostgre() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=example_db password=postgres sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	// use singular table name
	db.SingularTable(true)

	// auto migrate
	db.AutoMigrate(&model.User{})

	return db
}