package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var timeZoneLocation *time.Location

func init() {
	timeZoneLocation, _ = time.LoadLocation(os.Getenv("TZ"))
}

func main() {
	port := flag.Int("port", 3001, "Set port to listen on")
	flag.Parse()

	fmt.Printf("TCP server started on port: %d\n", *port)

	startServer(port)
}

func startServer(port *int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	log.Println("Connection made")
	log.Println(c.RemoteAddr())
	for {
		_, err := c.Write([]byte(time.Now().In(timeZoneLocation).Format("15:04:05 MST\n")))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
