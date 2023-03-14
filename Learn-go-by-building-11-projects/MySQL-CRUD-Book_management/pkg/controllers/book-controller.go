package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/DurianDan/go-bookstore/pkg/utils"
	"github.com/DurianDan/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= models.GetAllBooks()
	// turn Go Objects into JSON bytes
	// ultilize the field tags in Go struct
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request ) {
	vars:= mux.Vars(r)
	bookId := vars["boolId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	// ParseInt return both result and error 
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	// the db variable is not used, parsed it into _ 
	res, err := json.Marshal(bookDetails)
	if err != nil {
		fmt.Println("Cant make JSON from object %v",bookDetails)
	} else{
		w.Header().Set("Content-Type","pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	// create an empty Book struct 
	// asign the pointer to variable CreateBook
	utils.ParseBody(r, CreateBook)
	// UnMasharl the request r, parse it in to the struct Book ?? or interface CreateBook inside the struct Book ?
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId := vars["boolId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	// ParseInt return both result and error 
	if err != nil {
		fmt.Println("error while parsing")
	}

	book,_ := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
 
func UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := &models.Book{}
	utils.Parse(r,updateBook)
	vars:= mux.Vars(r) 
	bookId: vars
}