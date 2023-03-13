package models

import (
	"github.com/DurianDan/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"` 
	// contains 2 tags, empty tag gorm, and json tag name
	// gorm tag tell the database how to handle field Name
	// empty gorm tag means the DB will upload the data to column with the same name: "Name" 
	Author string `json:"author`
	Publication string `json:"publication"`
}

// tags are for encoding and decoding to/from JSON
// gorm.model provides basic fields to the struct "Book"

func init(){
	// initialize the database
	config.Connect()
	db = config.GetDB() // return the variable db, initiated inside the config package
	db.AutoMigrate(&Book{})
	// If table "Book" exist in db, it will update the schema (NOT THE DATA),
	// deleting/formating/adding columms
	// else, it just create a new table called Book
}

func (b *Book) CreateBook() *Book{
	// a function applied to the struct Book
	db.NewRecord(b)
	// check if b has existed in the db or not
	// needs to take in a pointer, to change the value b (by flaging it existed/not existed) 
	db.Create(&b) // upload the new record to database 
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	// return all records
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book 
	db := db.Where("ID=?".Id).Find(&getBook)
	return &getBook, db 
}

func DeleteBook(Id int64) Book {
	var book Book 
	db.Where("ID=?",ID).Delete(book)
	return book
}