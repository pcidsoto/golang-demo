package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pcidsoto/golang-demo/cmd/handlers"
	"github.com/pcidsoto/golang-demo/internal/book"
	"github.com/pcidsoto/golang-demo/models"
)

var DB *gorm.DB

func main() {
	server := gin.Default()

	database, error := gorm.Open("sqlite3", "test.db")

	if error != nil {
		fmt.Println("error: ", error)
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Book{})

	DB = database

	bookRepo := book.NewRespository(DB)
	bookServ := book.NewService(bookRepo)
	bookHandler := handlers.NewProduct(bookServ)

	server.GET("/books", bookHandler.GetAll())
	server.POST("/books", bookHandler.Store())

	server.Run(":8082")
}
