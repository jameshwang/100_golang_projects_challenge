package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/go-errors/errors"
)

var counterFileName = "id-counter.txt"

func init() {
	// TODO - this resets the counter each time the program restarts
	// check if file exists, otherwise, do the following
	err := ioutil.WriteFile(counterFileName, []byte("1"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Retrieve a data blob based on id
func Retrieve(id int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("%d.txt", id))
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// Create a data blob on the file system to store the string provided
func Create(data string) int {
	id := getNewID()
	fileName := fmt.Sprintf("%d.txt", id)
	err := ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func getNewID() int {
	data, err := ioutil.ReadFile(counterFileName)
	if err != nil {
		log.Fatal(err.(*errors.Error).ErrorStack())
	}

	currentID, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatal(err.(*errors.Error).ErrorStack())
	}

	nextID := currentID + 1

	err = ioutil.WriteFile(counterFileName, []byte(strconv.Itoa(nextID)), 0644)
	if err != nil {
		log.Fatal(err.(*errors.Error).ErrorStack())
	}
	return currentID
}
