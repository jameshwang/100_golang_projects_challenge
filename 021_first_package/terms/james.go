package terms

import "fmt"

type user struct {
	Name     string
	Age      int
	Address  string
	Employed bool
}

func init() {
	fmt.Println("the James package was loaded!")
}

// Return the name of the term
func Name() string {
	return "James Hwang"
}

func GetUser() user {
	james := user{Name: "James Wang", Age: 22, Address: "123 Main St.", Employed: true}
	return james
}

func (u user) ToString() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nAddress: %s\nEmployed: %t", u.Name, u.Age, u.Address, u.Employed)
}
