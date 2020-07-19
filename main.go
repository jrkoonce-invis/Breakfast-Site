package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jrkoonce-invis/BreakfastSite/Breakfast-Site/handlers"
	"github.com/sirupsen/logrus"
)

const (
	_port = ":9090"
	_wait = 10 * time.Second
)

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
		Addr:         _port,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.WithFields(logrus.Fields{
				"err": err,
			}).Fatal("Server failed to start")
		}

		log.Info(fmt.Sprintf("Server started on port %s", _port))
	}()

	// Create channel for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Recieve shutdown signals
	sig := <-stop
	ctx, cancel := context.WithTimeout(context.Background(), _wait)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.WithFields(logrus.Fields{
			"signal": sig,
			"err":    err,
		}).Warn("Error shutting down server")
	} else {
		log.WithFields(logrus.Fields{
			"signal": sig,
		}).Info("Server gracefully shutdown")
	}
}
