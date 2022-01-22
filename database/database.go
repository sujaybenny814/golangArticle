package database

import (
	"fmt"

	model "github.com/micro1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := "host=localhost user=postgres password='postgres' dbname=postgres port=5432 "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database ")
	}
	db = db.Exec("CREATE DATABASE article")
	if db.Error != nil {
		fmt.Println("Unable to create DB test_db, attempting to connect assuming it exists...")
		db, err = gorm.Open(postgres.Open("host=localhost user=postgres password='postgres' dbname=article port=5432 "), &gorm.Config{})
		if err != nil {
			fmt.Println("Unable to connect to article :", err)
			panic(err)
		}
	}
	db.AutoMigrate(&model.Article{})
	return db
}
