package main

import (
	"encoding/json"
	"net/http"
	"spider_quotes_api/quotes"
	"strings"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/quotes", getQuotesHandler).Methods("GET")
	r.HandleFunc("/quotes/search", searchQuotesHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getQuotesHandler(w http.ResponseWriter, r *http.Request) {
	quotesList := quotes.GetQuotes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotesList)
}

func searchQuotesHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "Keyword is required", http.StatusBadRequest)
		return
	}
	filteredQuotes := quotes.SearchQuotes(strings.ToLower(keyword))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredQuotes)
}