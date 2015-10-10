package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Make the commands and interface RESTful. so using web server
	// take commands to retrieve data by id
	// take commands to store data by id
	//
	// Example request will look like:
	// curl -XPOST db.onceprime.com/ -d "data='store this awesome info for me'"
	// curl -XGET db.onceprime.com/:id

	r := mux.NewRouter()
	r.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		fmt.Fprintln(w, Retrieve(id))
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		id := Create(string(body))
		fmt.Fprintln(w, id)
	}).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
