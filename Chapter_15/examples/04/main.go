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
		fmt.Println("The connection to the DB was successfully initialized")
	}
	updateStatement := `UPDATE test SET name = $1 where id = $2`

	updateResult, updateResultErr := db.Exec(updateStatement, "well", 2)
	if updateResultErr != nil {
		panic(updateResultErr)
	}

	updatedRecords, updatedRecordsErr := updateResult.RowsAffected()
	if updatedRecordsErr != nil {
		panic(updatedRecordsErr)
	}
	fmt.Println("Number of records updated: ", updatedRecords)
	db.Close()
}
