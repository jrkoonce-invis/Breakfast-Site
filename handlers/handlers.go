package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type requestManager struct {
	log *logrus.Logger
}

func CreateHandlers(log *logrus.Logger) *requestManager {
	return &requestManager{log}
}

// Handles the GET requests
func (manager *requestManager) GetItems(rw http.ResponseWriter, req *http.Request) {
	manager.log.WithFields(logrus.Fields{
		"request-type": "GET",
	}).Info("Handle: GET Request")
	rw.Write([]byte("These are the products!"))
}

// Handles the POST requests
func (manager *requestManager) PostItem(rw http.ResponseWriter, req *http.Request) {
	manager.log.WithFields(logrus.Fields{
		"request-type": "POST",
	}).Info("Handle: POST Request")
	rw.Write([]byte("You posted a product!"))
}

// Handles the PUT requests
func (manager *requestManager) PutItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	manager.log.WithFields(logrus.Fields{
		"request-type": "PUT",
	}).Info("Handle: PUT Request")
	rw.Write([]byte(fmt.Sprintf("You replaced a product: %s", id)))
}

// Handles the DELETE requests
func (manager *requestManager) DeleteItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	manager.log.WithFields(logrus.Fields{
		"request-type": "DELETE",
	}).Info("Handle: DELETE Request")
	rw.Write([]byte(fmt.Sprintf("You deleted a product: %s", id)))
}
