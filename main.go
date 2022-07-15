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

type Novel struct {
	ID     string  `json: "id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json: "author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var novels []Novel

func main() {
	r := mux.NewRouter()

	novels = append(novels, Novel{ID: "1", Isbn: "438227", Title: "The 300", Author: &Author{Firstname: "Arthur", Lastname: " Chandi"}})
	novels = append(novels, Novel{ID: "2", Isbn: "767780", Title: "The Twelve", Author: &Author{Firstname: " Maxwell", Lastname: "Uncle"}})

	r.HandleFunc("/novels", getNovels).Methods("GET")
	r.HandleFunc("/novels/{id}", getNovel).Methods("GET")
	r.HandleFunc("/novels", createNovel).Methods("POST")
	r.HandleFunc("/novels/{id}", updateNovel).Methods("PUT")
	r.HandleFunc("/novel/{id}", deleteNovel).Methods("DELETE")

	fmt.Printf("Starting server at port 9080\n")
	log.Fatal(http.ListenAndServe(":9080", r))
}

func createNovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var novel Novel
	_ = json.NewDecoder(r.Body).Decode(&novel)
	novel.ID = strconv.Itoa(rand.Intn(1000000000))
	novels = append(novels, novel)
	json.NewEncoder(w).Encode(novel)
}

func deleteNovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, element := range novels {
		if element.ID == params["id"] {
			novels = append(novels[:index], novels[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(novels)
}

func getNovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, element := range novels {
		if element.ID == params["id"] {
			json.NewEncoder(w).Encode(element)
			return
		}
	}
}

func getNovels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novels)
}

func updateNovel(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//get params
	params := mux.Vars(r)
	//loop over the novels range
	for index, element := range novels {
		if element.ID == params["id"] {
			//delete novel with the given id - deleteNovel()
			novels = append(novels[:index], novels[index+1:]...)
			//add new novel- createNovel()
			var novel Novel
			_ = json.NewDecoder(r.Body).Decode(&novel)
			novel.ID = params["id"]
			novels = append(novels, novel)
			json.NewEncoder(w).Encode(novel)
			return
		}

	}
}
