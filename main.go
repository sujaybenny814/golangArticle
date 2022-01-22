package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro1/database"
	hd "github.com/micro1/handlers"
)

func main() {
	DB := database.Init()
	h := hd.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/createArticle", h.AddArticle).Methods("POST")
	r.HandleFunc("/getAllArticle", h.GetAllArticle).Methods("GET")
	r.HandleFunc("/updateArticle/{id}", h.UpdateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
