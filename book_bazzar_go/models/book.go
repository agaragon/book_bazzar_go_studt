package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	PublishYear int    `json:"publish_year"`
	Description string `json:"description"`
}

// CreateBook creates a new book record in the database
func (b *Book) CreateBook() error {
	return DB.Create(b).Error
}

// GetBook retrieves a book by ID from the database
func (b *Book) GetBook(id uint) error {
	return DB.First(b, id).Error
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() ([]Book, error) {
	var books []Book
	err := DB.Find(&books).Error
	return books, err
}

// UpdateBook updates an existing book record in the database
func (b *Book) UpdateBook() error {
	return DB.Save(b).Error
}

// DeleteBook deletes a book record from the database
func (b *Book) DeleteBook() error {
	return DB.Delete(b).Error
}
