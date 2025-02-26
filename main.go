package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Server started on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func snippetCreate(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("Create a new snippet..."))
}

func snippetview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
