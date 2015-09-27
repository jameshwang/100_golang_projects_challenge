package main

import "fmt"

func main() {
	// Make the commands and interface RESTful. so using web server
	// take commands to retrieve data by id
	// take commands to store data by id
	//
	// Example request will look like:
	// curl -XPOST db.onceprime.com/ -d "data='store this awesome info for me'"
	// curl -XGET db.onceprime.com/:id

	Create("first time!\n")
	Create("second time!\n")
	Create("third time!\n")

	fmt.Printf(Retrieve(3))
}
