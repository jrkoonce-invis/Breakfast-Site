package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRouter)

	log.Fatal(http.ListenAndServe(":9090", nil))
	fmt.Println("Server running on port :9090")
}

func getRouter(rw http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://server:3000/items")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(rw, string(body))
	fmt.Fprint(rw, "YAYYY!!")
}
