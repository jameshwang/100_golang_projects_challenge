package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	done := make(chan int)
	go watchClock(3001)
	go watchClock(3002)
	// go watchClock(3003)
	// go watchClock(3004)

	<-done
}

func watchClock(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(message)
	}
}
