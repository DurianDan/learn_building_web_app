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
	Isbn string `json:"isbn"` // is like barcode for films
	Title string `json:"title"`
	Director *Director `json:"director"`
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
	var new_movie Movie 
	json.NewDecoder(r.Body).Decode(&new_movie)

	new_movie.ID = strconv.Itoa(rand.Intn(10000000000000))
	movies = append(movies, new_movie)
	json.NewEncoder(w).Encode(new_movie)
}

func updateMovie( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var update_movie Movie 
	request_update_idx := mux.Vars(r)["id"]

	// what is the  difference to params := mux.Vars(r) ?
	// mux.Vars(r) get the {id} at the end of the URL 
	// the below line of code get the Body of the request as a JSON
	json.NewDecoder(r.Body).Decode(&update_movie)
	update_movie.ID = request_update_idx

	for idx, item := range movies {
		if item.ID == request_update_idx {
			movies[idx] = update_movie
			json.NewEncoder(w).Encode(movies[idx])
			break
		}
	}

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
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at the port 8000: localhost:8000")

	log.Fatal(http.ListenAndServe(":8000",r))
}