package routes

import (
	"github.com/gorilla/mux"
	"github.com/DurianDan/go-bookstore/pkg/controllers"
)

func RegisterBookStore(router *mux.Router){
	// to initiate all routes automatically
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// the book route will be handle by the CreateBook function from the folder controllers
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}