package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models "book_bazzar_go/models"
	controllers "book_bazzar_go/controllers"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func setupRoutes(r *gin.Engine) {
    bookController := controllers.NewBookController()
    
    books := r.Group("/api/books")
    {
        books.POST("/", bookController.CreateBook)
        books.GET("/", bookController.GetAllBooks)
        books.GET("/:id", bookController.GetBook)
        books.PUT("/:id", bookController.UpdateBook)
        books.DELETE("/:id", bookController.DeleteBook)
    }
}

func main() {
	// Database connection
	models.DB = initDB()

	// Auto migrate the Book model
	err := models.DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	setupRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

