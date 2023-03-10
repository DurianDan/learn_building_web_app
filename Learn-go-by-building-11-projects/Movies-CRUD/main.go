package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"ID"` // json:"ID" is struct tag ??? later
	Isbn string `json: "isbn"` // is like barcode for films
	Title string `json:"title"`
	Director *Director `json:"director`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// set header of w, to return json type
	
	json.NewEncoder(w).Encode(movies)
	// return a slice of Movie structs
	// write it to a io.Writer object (which is http.ResponseWriter)
	// structs can be converted into json
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
}



func main(){
	r:=mux.NewRouter()

	movies = append(movies,
		Movie{"1","lskjdlkjf","Avatar",
			&Director{"Michael", "Jackson"}},
	)
	movies = append(movies,
		Movie{"2","lskjdl3k2","Avatar2",
			&Director{"Mikel", "Babson"}},
	)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// when client request (using GET method) to the route /movies...
	// execute the function getMovies(w mux.ResponseWriter, r *mux.Requests)
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}",delteMovies).Methods("DELETE")

	fmt.Println("Starting server at the port 8000: localhost:8000")

	log.Fatal(http.ListenAndServe(":8000",r))
}