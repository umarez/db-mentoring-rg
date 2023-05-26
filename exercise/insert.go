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
	var departments = []Department{
		{Name: "Department 1"},
		{Name: "Department 2"},
		{Name: "Department 3"},
		{Name: "Department 4"},
		{Name: "Department 5"},
	}
	db.Create(&departments)

	var employees = []Employee{
		{
			Name:          "Employee 1",
			Address:       "Jl. Kebon Jeruk",
			Age:           24,
			Birthdate:     "2004-01-01",
			Level:         "Staff",
			Id_department: 1,
		},
		{
			Name:          "Employee 2",
			Address:       "Jl. Kebon Jeruk",
			Age:           25,
			Birthdate:     "2001-01-01",
			Level:         "Staff",
			Id_department: 2,
		},
		{
			Name:          "Employee 3",
			Address:       "Jl. Kebon Jeruk",
			Age:           23,
			Birthdate:     "2003-01-01",
			Level:         "Staff",
			Id_department: 3,
		},
		{
			Name:          "Employee 4",
			Address:       "Jl. Kebon Jeruk",
			Age:           22,
			Birthdate:     "2005-01-01",
			Level:         "Staff",
			Id_department: 4,
		},
		{
			Name:          "Employee 5",
			Address:       "Jl. Kebon Jeruk",
			Age:           21,
			Birthdate:     "2006-01-01",
			Level:         "Staff",
			Id_department: 5,
		},
	}

	db.Create(&employees)

	fmt.Println("Insertion Success")
}
