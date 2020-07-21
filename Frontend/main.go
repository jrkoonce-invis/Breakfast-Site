package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	frontend := http.FileServer(http.Dir("./src"))
	http.Handle("/", frontend)

	fmt.Println("Server running on port :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

/*
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
}*/
