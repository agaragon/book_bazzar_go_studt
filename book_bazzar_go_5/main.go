package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"net/http"
)

var DB *gorm.DB

type Book struct {
	Id uint `json"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
}

func (book *Book) Create() error {
	return DB.Create(book).Error
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := book.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &book)	
	return
}

func initDB() *gorm.DB {
	dns := "host=localhost user=postgres password=postgres port=5432 dbname=book_bazzar_5"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to start database", err)
	}
	return db
}

func setupHandlers(e *gin.Engine) {
	books := e.Group("/api/books")

	books.POST("", createBookHandler)
}

func main() {
	DB = initDB()


	DB.AutoMigrate(&Book{})
	// Create routes engine
	r := gin.Default()

	setupHandlers(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("unable to start server", err)
	}
}