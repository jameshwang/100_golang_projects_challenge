package main

import (
	"encoding/json"
	"fmt"
)

type Model struct {
	UserID  uint   `json:"user_id"`
	Address string `json:"address,omitempty"`
}

func main() {
	model := Model{UserID: 23}
	b, _ := json.Marshal(model)
	fmt.Println(string(b))
}
