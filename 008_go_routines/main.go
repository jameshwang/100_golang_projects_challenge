package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	name string
	args string
}

func worker(work <-chan Message, finished chan<- string) {
	for {
		message := <-work
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		finished <- strings.Join([]string{message.name, "Done"}, " - ")
	}
}

func manager() {
	jobs := make(chan Message, 2000)
	finished := make(chan string, 5000)
	for i := 0; i < 2000; i++ {
		fmt.Println("creating worker: ", i)
		go worker(jobs, finished)
	}

	// work := []string{"clean kitchen", "paint wall", "take out garbage"}
	for i := 0; i < 2000; i++ {
		jobs <- Message{strconv.Itoa(i), "with force"}

	}

	for {
		message := <-finished
		fmt.Println("message: ", message)
	}
}

func main() {
	go manager()

	time.Sleep(3 * time.Second)
}
