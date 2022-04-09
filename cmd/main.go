package main

import (
	"github.com/devcarlosl/book-api/pkg/books"
	"github.com/devcarlosl/book-api/pkg/common/config"
	"github.com/devcarlosl/book-api/pkg/common/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	h := db.Init(&c)

	books.RegisterRoutes(r, h)

	r.Run(c.Port)
}
