package main

import (
	"fmt"

	"github.com/jameshwang/100_go_projects/021_first_package/terms"
)

func main() {
	fmt.Println(terms.Name())

	// arr := [...]int{34, 23, 65}
	arr := [...]string{"Red", "Blue", "Green", "Yellow"}

	for _, v := range arr[1:2] {
		fmt.Println(v)
	}

	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice = append(newSlice, 42)

	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[1] = 35

	fmt.Println(slice)
	fmt.Println(newSlice)

	fmt.Println(terms.GetUser().ToString())
}
