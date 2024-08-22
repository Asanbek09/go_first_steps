package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := "/static/index.html"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File %s does not exist", path)
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
