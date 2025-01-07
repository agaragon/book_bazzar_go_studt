package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"net/http"
	"log"
)

var DB *gorm.DB

type Book struct {
	Id uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
}

func (book *Book) CreateBook() error {
	return DB.Create(book).Error
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := book.CreateBook(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func setupHandlers(r *gin.Engine) {
	books := r.Group("/api/books")

	books.POST("", createBookHandler)
}

func initDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres port=5432 dbname=book_bazzar_3"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to start db:", err)
	}

	return db
}

func main() {
	r := gin.Default()

	DB = initDb()

	DB.AutoMigrate(&Book{})

	setupHandlers(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start err:", err.Error())
	}
}