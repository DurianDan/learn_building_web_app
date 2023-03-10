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

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // get params from request
	for idx, item := range movies {
		// ragne return the index and a copy of each value in a slice/array
		// each item is of struct "Movie"
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			// because movies[idx+1:] contains many elements
			// append() can only recieve a slice, and element, not something containing multiple elements 
			// in need to add ... to the end of it, to unwrap it into multiplt elements 
			break
			// this method can not work if in need of deleting multiple movies
		}
	}

	// return the remainning movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	get_id := mux.Vars(r)["id"]
	for _, item := range movies {
		if item.ID == get_id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	var new_movie = Movie{
		params["ID"],
		params["isbn"],
		params["title"],
		&Director{
			params["firstname"],
			params["lasstname"],
		},
	}
	movies = append(movies, new_movie)
	json.NewEncoder(w).Encode(new_movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	var update_idx int
	for idx, item := range movies {
		if item.ID == params["ID"]{
			update_idx = idx
			movies[idx] = Movie{
				params["ID"],
				params["isbn"],
				params["title"],
				&Director{
					params["firstname"],
					params["lastname"],
				},
			}
			break
		}
	}
	json.NewEncoder(w).Encode(movies[update_idx])
}

func main() {
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
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at the port 8000: localhost:8000")

	log.Fatal(http.ListenAndServe(":8000",r))
}