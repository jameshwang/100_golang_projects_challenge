package main

import (
	"fmt"
	"time"
)

func generator(c chan<- int) {
	for i := 1; i <= 20; i++ {
		c <- i
		time.Sleep(time.Millisecond * 200)
	}
	close(c)
}

func multiplier(input <-chan int, output chan<- int) {

	for value := range input {
		output <- value * value
	}

	close(output)
}

func printer(input <-chan int) {
	for value := range input {
		fmt.Println(value)
	}
}

func main() {
	start := make(chan int)
	multiply := make(chan int)

	go generator(start)
	go multiplier(start, multiply)
	printer(multiply)
}
