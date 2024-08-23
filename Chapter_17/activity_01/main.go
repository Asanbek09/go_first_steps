package main

import (
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// send the get request
	r, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	//get data from the response body
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	//return the response
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}
