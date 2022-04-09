package db

import (
	"fmt"
	"github.com/devcarlosl/book-api/pkg/common/config"
	"github.com/devcarlosl/book-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf(c.DBUrl)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
