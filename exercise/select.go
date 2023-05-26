package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID            uint
	Name          string
	Address       string
	Age           uint
	Birthdate     string
	Level         string
	Id_department uint
}

type Department struct {
	ID   uint
	Name string
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=test_db_camp port=5430 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	var employees []Employee

	db.Find(&employees)

	for _, employee := range employees {
		fmt.Println(employee)
	}
	var departments []Department

	db.Find(&departments)

	for _, department := range departments {
		fmt.Println(department)
	}
}
