package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var conn net.Conn

func init() {
	var err error
	conn, err = net.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Panic("Error while trying to connect to port 6379")
	}
}

func testRedisConnection() {
	fmt.Fprintf(conn, "PING\r\n")
	result, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	if result != "+PONG\r\n" {
		log.Panic("Redis db didn't respond correctly. Exiting...")
	}
}

func listenToRedisResults() {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	testRedisConnection()
	go listenToRedisResults()
	fmt.Fprintf(conn, "KEYS *\r\n")
	fmt.Fprintf(conn, "PING\r\n")
	time.Sleep(time.Second)
}
