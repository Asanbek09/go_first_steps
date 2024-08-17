package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=4321 host=127.0.0.1 port=5432 dbname=go_example sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("the connection to the DB was successfully initialized")
	}

	DBCreate := `CREATE TABLE public.test (
	id integer,
	name character varying COLLATE pg_catalog."default")
	WITH (
		OIDS = FALSE
	)`

	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was successfully created")
	}
	db.Close()
}
