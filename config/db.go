package config

import (
	"fmt"
	"go_web_project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "go_db"
	DBHost     = "localhost"
	DBPort     = "5432"
)

func DBInit() *gorm.DB {
	dbURL := GetPostgresConnectionString()
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}
