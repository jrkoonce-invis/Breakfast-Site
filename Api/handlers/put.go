package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jrkoonce-invis/Breakfast-Site/Api/data"
	"github.com/sirupsen/logrus"
)

// Handles the PUT requests
func (manager *requestManager) PutItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	manager.log.WithFields(logrus.Fields{
		"request-type": "PUT",
	}).Info("Handle: PUT Request")

	item, err := ioutil.ReadAll(req.Body)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error reading request body")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	uintID, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error converting string to uint16")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ReplaceDatum(item, uint16(uintID))
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error unmarshalling json or (more often) error validating request body")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write([]byte(fmt.Sprintf("You replaced a product: %s", id)))
}
