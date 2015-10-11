package hello

import "testing"

// func TestHello(t *testing.T) {
// 	if Hello() != "hello" {
// 		t.Errorf("Expected 'hello', but got '%s'", Hello())
// 	}
// }

func TestGreeting(t *testing.T) {
	james := Person{Name: "James"}
	if james.greeting() != "James" {
		t.Errorf("Expected 'asdf', but got '%s'", james.greeting())
	}
}
