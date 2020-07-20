package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jrkoonce-invis/BreakfastSite/Breakfast-Site/data"
	"github.com/sirupsen/logrus"
)

// Handles the DELETE requests
func (manager *requestManager) DeleteItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	manager.log.WithFields(logrus.Fields{
		"request-type": "DELETE",
	}).Info("Handle: DELETE Request")

	uintID, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		manager.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error converting string to uint16")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	data.DeleteDatum(uint16(uintID))

	rw.Write([]byte(fmt.Sprintf("You deleted a product: %s", id)))
}
