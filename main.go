package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie struct
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

// Director struct
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init movies var as a slice Movie struct
var movies []Movie


// Get all movies - endpoint handler function
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


// Get single movie - endpoint handler function
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	// loop through the movies and find the movie with the ID
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// if the movie is not found, return an empty movie
	json.NewEncoder(w).Encode(&Movie{})
}


// Create a new movie - endpoint handler function
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


// Update a movie - endpoint handler function
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}


// Delete a movie - endpoint handler function
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}







func main() {
	// Init Router
	r := mux.NewRouter()

	// Hardcoded data/Mock data - can be replaced with a database later
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "847564", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// handle the routes/endpoints for the API
	r.HandleFunc("/api/movies", getMovies).Methods("GET") // get all movies
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET") // get single movie
	r.HandleFunc("/api/movies", createMovie).Methods("POST") // create a movie
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT") // update a movie
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE") // delete a movie

	// print the message to the console to show the server is running on port 8000
	fmt.Printf("Starting server on port 8000\n")

	// log the error if the server fails to start
	log.Fatal(http.ListenAndServe(":8000", r))
}