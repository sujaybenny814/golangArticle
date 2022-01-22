package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/micro1/models"
)

func (h handler) AddArticle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to create an article")
	}
	var article model.Article
	json.Unmarshal(body, &article)

	if response := h.DB.Create(&article); response.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to create an article")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Created an article")
	}
}

func (h handler) GetAllArticle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var articles []model.Article
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")
	// search = "^" + search

	if response := h.DB.Limit(limit).Offset(page).Order("title").Where("title = ?", search).
		Find(&articles); response.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("No Article found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(articles)
	}
}

func (h handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if response := h.DB.Model(&model.Article{}).Where("id = ?", id).Update("status", "approved"); response.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Failed to Update Article")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Update successfully")
	}
}
