package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jrkoonce-invis/BreakfastSite/Breakfast-Site/handlers"
)

const port = ":9090"

func main() {
	// Create the gorilla servemux
	mux := mux.NewRouter()

	// Create a subrouter specifically for the REST api calls
	itemsManager := mux.PathPrefix("/items").Subrouter()

	// Setting up the REST handlers
	itemsManager.HandleFunc("", handlers.GetItems).Methods("GET")
	itemsManager.HandleFunc("", handlers.PostItem).Methods("POST")
	itemsManager.HandleFunc("/{id}", handlers.PutItem).Methods("PUT")
	itemsManager.HandleFunc("/{id}", handlers.DeleteItem).Methods("DELETE")

	// Create a new server
	server := http.Server{
		Addr:         port,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
		os.Exit(1)
	}

	fmt.Println("Starting server on port", port)

}
