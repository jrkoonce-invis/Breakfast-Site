package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetItems(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Handle: GET Request")
	rw.Write([]byte("These are the products!"))
}

func PostItem(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Handle: POST Request")
	rw.Write([]byte("You posted a product!"))
}

func PutItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	fmt.Println("Handle: PUT Request")
	rw.Write([]byte(fmt.Sprintf("You replaced a product: %s", id)))
}

func DeleteItem(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	fmt.Println("Handle: DELETE Request")
	rw.Write([]byte(fmt.Sprintf("You deleted a product: %s", id)))
}
