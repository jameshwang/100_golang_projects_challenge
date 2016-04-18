package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Customer struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
}

func (c *Customer) String() string {
	return fmt.Sprintf("%s %s: %s", c.Firstname, c.Lastname, c.Email)
}

type Order struct {
	gorm.Model
	Product    Product
	ProductID  int
	Customer   Customer
	CustomerID int
}

type Product struct {
	gorm.Model
	Name     string
	Count    int
	Vendor   Vendor
	VendorID int
}

type Vendor struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("postgres", "user=jameshwang dbname=jameshwang sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	db.AutoMigrate(&Customer{}, &Order{}, &Product{}, &Vendor{})

	var customer Customer
	var order Order

	db.Last(&order)

	println(db.Model(&order).Related(&customer))
	println(customer.String())

	// db.Create(&Order{
	// 	Product: Product{
	// 		Name:  "iPhone",
	// 		Count: 2,
	// 		Vendor: Vendor{
	// 			Name: "Apple",
	// 		},
	// 	},
	// 	Customer: Customer{
	// 		Firstname: "Jimmy",
	// 		Lastname:  "Hollywood",
	// 		Email:     "jimmy@onceprime.com",
	// 	},
	// })
}
