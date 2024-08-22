package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

var s string

func main() {
	t, _ := template.New("mytemplate").Parse(s)
	http.HandleFunc("/http1", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Text string
		}{
			Text: "Hello, there",
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8085", nil))
}
