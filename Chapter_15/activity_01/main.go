package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// import _ <package_name> is a special import statement that tells Go to import a package solely for its side effects

func main() {
	db, err := sql.Open("postgres", "user=postgres password=4321 host=127.0.0.1 port=5432 dbname=go_example sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized")
	}

	connectivity := db.Ping()
	if connectivity != nil {
		panic(err)
	} else {
		fmt.Println("Good to go!")
	}

	defer db.Close()
}
