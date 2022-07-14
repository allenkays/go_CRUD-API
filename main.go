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
	Id     string  `json: "Id"`
	Isbn   string  `json:"Isbn"`
	title  string  `json:"title"`
	Author *Author `json: "Author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var novels []Novel

func main() {
	r := mux.NewRouter()

	novels = append(novels, Novel{Id: "1", Isbn: "438227", title: "The 300", Author: &Author{Firstname: "Arthur", Lastname: " Chandi"}})
	novels = append(novels, Novel{Id: "2", Isbn: "767780", title: "The Twelve", Author: &Author{Firstname: " Maxwell", Lastname: "Uncle"}})

	r.HandleFunc("/novels", getNovels).Methods("GET")
	r.HandleFunc("/novels/{Id}", getNovel).Methods("GET")
	r.HandleFunc("/novels", createNovel).Methods("POST")
	r.HandleFunc("/novels/{Id}", updateNovel).Methods("PUT")
	r.HandleFunc("/novel/{Id}", deleteNovel).Methods("DELETE")

	fmt.Printf("Starting server at port 9080\n")
	log.Fatal(http.ListenAndServe(":9080", r))
}

func createNovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var novel Novel
	_ = json.NewDecoder(r.Body).Decode(novel)
	novel.Id = strconv.Itoa(rand.Intn(1000000000))
	novels = append(novels, novel)
	json.NewEncoder(w).Encode(novel)
}

func deleteNovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, element := range novels {
		if element.Id == params["Id"] {
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
		if element.Id == params["Id"] {
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, element := range novels {
		if element.Id == params["Id"] {
			//deleteNovel()
			novels = append(novels[:index], novels[index+1:]...)
			//createNovel()
			var novel Novel
			_ = json.NewDecoder(r.Body).Decode(&novel)
			novel.Id = params["Id"]
			novels = append(novels, novel)
			json.NewEncoder(w).Encode(novel)
			return
		}

	}
}
