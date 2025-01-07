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
	ISBN string `json:"isbn"`
	PublishYear int `json:"publish_year"`
}

func InitDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres port=5432 dbname=book_bazzar_2"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func (b *Book) CreateBook() error {
	return DB.Create(b).Error
}

func CreateBookHandler(c *gin.Context) {
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

func setupRoutes(r *gin.Engine) {
	books := r.Group("/api/books")

	books.POST("", CreateBookHandler)
}

func main() {
	r := gin.Default()

	DB = InitDb()
	DB.AutoMigrate(&Book{})

	setupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}