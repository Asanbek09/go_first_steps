package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Получаем текущий рабочий каталог
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Определяем путь к index.html, который находится рядом с main.go
	filePath := filepath.Join(wd, "index.html")
	fmt.Println(filePath)

	// Обслуживаем index.html по корневому пути "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})

	// Обслуживаем статические файлы из папки "stylish" по пути "/statics/"
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir(filepath.Join(wd, "public")))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
