package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type Book struct {
	Id uint `json:"id" gorm:"primary_key"` 
	Title string `json:"title"`
	Author string `json:"author"`
}

func (book *Book) Create() error {
	return DB.Create(&book).Error
}

func initDB() *gorm.DB{
	dsn := "user=postgres password=postgres port=5432 host=localhost dbname=book_bazzar_4"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to start db", err.Error())
	}
	return db
}

func CreateBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := book.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func setupHandlers(e *gin.Engine) {
	books := e.Group("/api/books")

	books.POST("", CreateBookHandler)
}

func main(){

	r := gin.Default()

	DB = initDB()

	DB.AutoMigrate(&Book{})

	setupHandlers(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("unable to start server: ", err)
	}
}