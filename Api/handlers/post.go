package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/jrkoonce-invis/Breakfast-Site/Api/data"
	"github.com/sirupsen/logrus"
)

// Handles the POST requests
func (manager *requestManager) PostItem(rw http.ResponseWriter, req *http.Request) {
	manager.log.WithFields(logrus.Fields{
		"request-type": "POST",
	}).Info("Handle: POST Request")

	item, err := ioutil.ReadAll(req.Body)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error marshalling json or (more often) error validating request body")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.PostData(item)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error unmarshalling json")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Product Posted Successfully"))
}
