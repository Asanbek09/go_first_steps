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

	defer db.Close()

	emptyTable, emptyTableErr := db.Exec("truncate table number")
	if emptyTableErr != nil {
		panic(emptyTableErr)
	}

	rowsAffected, err := emptyTable.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows affected by truncate: %d\n", rowsAffected)

	dropTable, dropTableErr := db.Exec("drop table number")
	if dropTableErr != nil {
		panic(dropTableErr)
	}
	rowsAffected, err = dropTable.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Table dropped, rows affected: %d\n", rowsAffected)
}
