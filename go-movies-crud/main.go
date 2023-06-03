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

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(500))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
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
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "101", Title: "ABCD", Director: &Director{Firstname: "Prabhu ", Lastname: "Deva"}})
	movies = append(movies, Movie{ID: "2", Isbn: "202", Title: "ABCD2", Director: &Director{Firstname: "Prabhu", Lastname: "Deva"}})
	movies = append(movies, Movie{ID: "3", Isbn: "303", Title: "RACE", Director: &Director{Firstname: "Rohit", Lastname: "Shetty"}})
	movies = append(movies, Movie{ID: "4", Isbn: "404", Title: "RACE2", Director: &Director{Firstname: "Rohit", Lastname: "Shetty"}})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deletemovie).Methods("DELETE")
	fmt.Printf("Starting the port :8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
