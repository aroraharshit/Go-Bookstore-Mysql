package models

import (
	"github.com/harshit/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
	"fmt"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(100)" json:"name"`
	Author      string `gorm:"column:author;type:varchar(100)" json:"author"`
	Publication string `gorm:"column:publication;type:varchar(100)" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	fmt.Printf("%+v\n", Books) 
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
