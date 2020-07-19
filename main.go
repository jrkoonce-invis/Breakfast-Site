package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jrkoonce-invis/BreakfastSite/Breakfast-Site/handlers"
	"github.com/sirupsen/logrus"
)

const port = ":9090"

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	// Create the gorilla servemux and universal logger (sirupsen/logrus)
	mux := mux.NewRouter()
	log := logrus.New()

	// Create the handlers
	requestManager := handlers.CreateHandlers(log)

	// Create a subrouter specifically for the REST api calls and register them
	itemsManager := mux.PathPrefix("/items").Subrouter()

	itemsManager.HandleFunc("", requestManager.GetItems).Methods("GET")
	itemsManager.HandleFunc("", requestManager.PostItem).Methods("POST")
	itemsManager.HandleFunc("/{id}", requestManager.PutItem).Methods("PUT")
	itemsManager.HandleFunc("/{id}", requestManager.DeleteItem).Methods("DELETE")

	// Create a new server
	server := http.Server{
		Addr:         port,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Server failed to start")
	}

	log.Info(fmt.Sprintf("Server started on port %s", port))
}
