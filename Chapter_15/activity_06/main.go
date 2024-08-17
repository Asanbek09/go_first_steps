package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

func main() {
	connection_string := "user=postgres password=4321 host=127.0.0.1 port=5432 dbname=go_example sslmode=disable"
	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&User{})

	//u := &User{FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com"}
	//db.Create(u)
	//db.Create(&User{FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com"})
	//db.Create(&User{FirstName: "James", LastName: "Smith", Email: "james.smith@gmail.com"})
	//var user User
	//db.First(&user, 1)
	//db.First(&user, "last_name = ?", "Doe")
	///db.First(&user, "last_name = ? and first_name = ? ", "Smith", "James")
	//db.First(&user, &User{FirstName: "James", LastName: "Smith"})

	var users []User
	//db.Find(&users, &User{LastName: "Smith"})

	tx := db.Find(&users, &User{LastName: "Smith"})
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}
