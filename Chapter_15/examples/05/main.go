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
	deleteStatement := `delete from test where id = $1`

	deleteResult, deleteResultErr := db.Exec(deleteStatement, 2)
	if deleteResultErr != nil {
		panic(deleteResultErr)
	}

	deletedRecords, deletedRecordsErr := deleteResult.RowsAffected()
	if deletedRecordsErr != nil {
		panic(deletedRecordsErr)
	}
	fmt.Println("Number of records deleted: ", deletedRecords)
	db.Close()
}
