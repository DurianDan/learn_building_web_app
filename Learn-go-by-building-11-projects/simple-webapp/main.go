package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		// write the formated string to the writer object w
		return
	}
	fmt.Fprintf(w, "POST request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	// when pass a function into http.HandleFunc()
	// the function will recieve a tool to Write a response: ResponseWriter
	// and a pointer to the request of the client *http.Request
	// ==> in Go, you can only edit a variable through a pointer 
	if r.URL.Path != "/hello" {
		// check if the hellHandler is handling requests to the route "/hello"
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return 
	}
	if r.Method != "GET"{
		// a request from a browser is GET by default
		http.Error(w, "method is not supported", http.StatusNotFound)
		return		
	}
	fmt.Fprintf(w,"Hello! This is an app made by Huy")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	
	http.Handle("/",fileServer) 
	// handle the highest route of the website
	// using a single file: index.html
	http.HandleFunc("/form", formHandler)
	// handle the route /form with a custom function
	http.HandleFunc("/hello", helloHandler)
	// handle a custom function

	fmt.Println("Start server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err!=nil{
		// create a server
		// this function returns an error if unsuccess or nil if success
		log.Fatal(err)
	}
}
