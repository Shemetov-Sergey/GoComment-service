package db

import (
	"log"

	"github.com/Shemetov-Sergey/GoComment-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Comment{}, &models.CommentChild{})

	return Handler{db}
}
