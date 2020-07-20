package handlers

import (
	"net/http"

	"github.com/jrkoonce-invis/Breakfast-Site/Api/data"
	"github.com/sirupsen/logrus"
)

// Handles the GET requests
func (manager *requestManager) GetItems(rw http.ResponseWriter, req *http.Request) {
	manager.log.WithFields(logrus.Fields{
		"request-type": "GET",
	}).Info("Handle: GET Request")

	items, err := data.GetData(manager.log)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error marshalling json")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(items))
}
