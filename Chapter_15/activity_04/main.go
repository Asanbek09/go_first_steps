package main

import (
	"database/sql"
	"fmt"
	"math/big"

	_ "github.com/lib/pq"
)

func main() {
	var number int64
	var prop string
	var primeSum int64
	var newNumber int64

	db, err := sql.Open("postgres", "user=postgres password=4321 host=127.0.0.1 port=5432 dbname=go_example sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized")
	}
	allTheNumbers := "SELECT * FROM Number"
	numbers, err := db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	primeSum = 0
	result, err := numbers.Query()
	fmt.Println("The list of prime numbers: ")
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Print(" ", number)
		}
	}

	err = numbers.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("\n The total sum of prime numbers in this range is : ", primeSum)
	remove := "delete from number where property = $1"
	removeResult, err := db.Exec(remove, "Even")
	if err != nil {
		panic(err)
	}
	modifiedRecords, err := removeResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("The number of rows removed: ", modifiedRecords)
	fmt.Println("Updating numbers ...")
	update := "update number set nuumber = $1 where nuumber=$2 and property=$3"
	allTheNumbers = "select * from number"
	numbers, err = db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	result, err = numbers.Query()
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		newNumber = number + primeSum
		_, err = db.Exec(update, newNumber, number, prop)
		if err != nil {
			panic(err)
		}
	}
	numbers.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("The execution is now complete...")
	db.Close()
}
