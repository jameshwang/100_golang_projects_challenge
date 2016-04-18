package main

import "fmt"

type Jim struct {
	name string
}

func (j Jim) Write(w []byte) (int, error) {
	fmt.Printf("Printing from jim: %s\n", w)
	return 5, nil
}

func (j Jim) String() string {
	return "Hello Jim"
}

func main() {
	var j Jim
	fmt.Println("hello!")
	fmt.Fprintf(j, "asdf")
	fmt.Println(j)
}
