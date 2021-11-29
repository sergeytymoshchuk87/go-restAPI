package main

import (
	"gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	get := r.PathPrefix("/get").Subrouter()
	get.HandleFunc("/all", ShowAll).Methods(http.MethodGet)
	get.HandleFunc("/{id}", GetByID).Methods(http.MethodGet)
	get.HandleFunc("/name/{name}", GetByName).Methods(http.MethodGet)

	post := r.PathPrefix("/post").Subrouter()
	post.HandleFunc("/new", AddNew).Methods(http.MethodPost)

	r.HandleFunc("/delete", Remove).Methods(http.MethodDelete)

	update := r.PathPrefix("/update/{id}").Subrouter()

	// Making room for unknown peculiarities in systems by allowing POST and PATCH (I don't really know, i did this out of intuition)
	update.HandleFunc("/name", UpdateName).Methods(http.MethodPatch, http.MethodPost)
	update.HandleFunc("/price", UpdatePrice).Methods(http.MethodPatch, http.MethodPost)
	update.HandleFunc("/time", UpdateTime).Methods(http.MethodPatch, http.MethodPost)
	update.HandleFunc("/all", UpdateAll).Methods(http.MethodPatch, http.MethodPost)
	http.ListenAndServe(":8080", r)

}
