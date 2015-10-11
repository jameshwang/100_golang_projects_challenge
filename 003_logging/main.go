package main

import (
	"fmt"
	"runtime"
)

func main() {
	// logfile, _ := os.Create("./log.txt")
	// defer logfile.Close()
	//
	// logger := log.New(logfile, "Example ", log.LstdFlags|log.Lshortfile)
	// logger.Println("This is a regular message.")
	// logger.Fatalln("This is a fatal error.")
	// logger.Println("This is the end of the function.")

	// conn, err := net.Dial("tcp", "localhost:1902")
	// if err != nil {
	// 	panic("Failed to connect to localhost:1902")
	// }
	// defer conn.Close()
	//
	// f := log.Ldate | log.Lshortfile
	// logger := log.New(conn, "example ", f)
	//
	// logger.Println("This is a regular message.")
	// logger.Fatalln("This is a panic.")

	// priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	// flags := log.Ldate | log.Lshortfile
	// logger, err := syslog.NewLogger(priority, flags)
	// if err != nil {
	// 	fmt.Printf("Can't attach to syslog: %s", err)
	// 	return
	// }
	// logger.Println("This is a test message.")

	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)
}
