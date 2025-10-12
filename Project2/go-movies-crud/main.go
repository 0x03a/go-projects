package main

import (
	"encoding/json" // used to convert data to and from JSON format (encoding/decoding)
	"fmt"
	"log"       // for error logging
	"math/rand" // for generating random IDs
	"net/http"  // for creating HTTP server and handling requests
	"strconv"   // to convert numbers to strings

	"github.com/gorilla/mux" // for routing requests to the correct handler
)

// note: we will not be using any database for this project.
// we will use struct and slice to store our data temporarily in memory.

// --------------------------------------------------------------------------
// 🧱 1. Struct
// A struct in Go is like a data model — it defines the structure of your data (like a table in a database).

// 📋 2. Slice
// A slice is a dynamic list that can hold multiple structs.
// Example: var users []User  → creates a list to store many User entries.
// A slice in Go is like a dynamic list or array.
// // You can store multiple User entries in a slice.
// //Example: //var users []User //This creates an empty list to hold multiple user
// --------------------------------------------------------------------------

// 💡 What is JSON?
// JSON (JavaScript Object Notation) is a lightweight data format used for sending and receiving data between
// a client (like Postman or a frontend app) and a server (like this Go backend).
// Example JSON data looks like this:
// {
//   "id": "1",
//   "title": "Inception",
//   "isbn": "12345",
//   "director": {
//       "firstname": "Christopher",
//       "lastname": "Nolan"
//   }
// }

// 💬 Why do we use JSON here?
// Because when we send or receive data through Postman or an API request, it comes in JSON format.
// Go needs to know how to convert (encode/decode) that data correctly.
// The `json:"fieldname"` tags tell Go what each field should be called in JSON.

// -------------------------REAL CODE BEGINS------------------------------

// Movie Struct (each movie has one director)
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// So when we send or receive data using Postman,
// we’ll use lowercase field names like id, isbn, title, and director.
// These json tags make sure Go knows how to properly encode and decode (When we use Go’s encoding/json package: Encoding means 👉 converting Go data → JSON format
//Decoding means 👉 converting JSON format → Go data)
// the data between Go structs and JSON format.
//----------------

// now, we will define a variable called movies (it will be a slice of the type movie)
var movies []Movie // dynamic list that can hold multiple movie of type struct

// getMovies handles the "GET /movies" request.
// It sends back all the movies in JSON format.
func getMovies(w http.ResponseWriter, r *http.Request) {

	// 'w' (http.ResponseWriter) is used to send a response back to the client (like Postman or browser).
	// Whatever you write to 'w' is what the client receives.

	// 'r' (*http.Request) represents the incoming HTTP request.
	// It contains details like headers, URL, method (GET/POST), and body data sent by the client.

	// Set the response header to tell the client that the data being sent is in JSON format.
	w.Header().Set("Content-Type", "application/json") // tells the client that the server’s response body is JSON data

	// Encode the 'movies' slice (our in-memory data) into JSON format
	// and send it in the response using 'w'.
	// Basically, this converts Go structs → JSON → sends to the client.
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
	// after removing returning existing movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies { // `_` this is blank indentifier
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// steps:
	// set json content-type
	// params
	// loop Over the movies,range
	// delete the movie with the id, that we have sent (as this way is not good, but as we are not working with database we have to do this.)
	// add a new movie - the movie that we send in the body of postman

	w.Header().Set("Content-Type", "application/json") // set json content-type
	params := mux.Vars(r)                              // params
	// loop Over the movies,range
	for index, item := range movies {
		if item.ID == params["id"] {
			// deleted
			movies = append(movies[:index], movies[index+1:]...)
			// now addung
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}
func main() {

	// Create a new router using the Gorilla Mux package.
	// A router helps direct incoming HTTP requests to the correct handler functions.
	r := mux.NewRouter() // Handles _dynamic routes_ for APIs or pages

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie1", Director: &Director{Firstname: "john", Lastname: "elia"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie2", Director: &Director{Firstname: "nhoj", Lastname: "aile"}})

	// Define routes (endpoints) for your API and link them with their handler functions.

	// This route handles GET requests to "/movies"
	// It calls the getMovies function, which will return a list of all movies.
	r.HandleFunc("/movies", getMovies).Methods("GET")

	// This route handles GET requests to "/movies/{id}"
	// It calls getMovie, which returns details of a single movie by its ID.
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	// This route handles POST requests to "/movies"
	// It calls createMovie, which adds a new movie to the list.
	r.HandleFunc("/movies", createMovie).Methods("POST")

	// This route handles PUT requests to "/movies/{id}"
	// It calls updateMovie, which updates details of an existing movie.
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	// This route handles DELETE requests to "/movies/{id}"
	// It calls deleteMovie, which removes a movie from the list.
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Print a message to the terminal indicating the server is starting.
	fmt.Printf("Starting server at port 8000\n")

	// Start the HTTP server on port 8000 and attach the router (r) to it.
	// log.Fatal ensures that if the server fails to start, the error is logged and the program stops.
	log.Fatal(http.ListenAndServe(":8000", r))

}
