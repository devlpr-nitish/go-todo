package db

import (
	"log"

	"github.com/devlpr-nitish/todo/internal/config"
	"github.com/devlpr-nitish/todo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func Connect(cfg config.Config) *gorm.DB{

	db, err := gorm.Open(postgres.Open(cfg.DbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Auto migrate table
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil{
		log.Fatal("failed to migrate database: ", err)
	}

	return db
}