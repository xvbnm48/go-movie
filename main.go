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
	ID string `json:"id"`	
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firtsname string `json:"firtsname"`	
	Lastname string `json:"lastname"`

}

var movies []Movie

func main(){
	r := mux.NewRouter()

	movies =append(movies, Movie{ID: "48", Title: "movie 1", Director: &Director{Firtsname: "makoto", Lastname: "shinkai"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies" createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server on port 8080\n")	
	log.Fatal(http.ListenAndServe(":8080", r))
}