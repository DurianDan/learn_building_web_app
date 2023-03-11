package test

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json:"ID"`
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies/{id}", updateMovies).Method("PUT")

	http.ListenAndServe(":8000",r)
}