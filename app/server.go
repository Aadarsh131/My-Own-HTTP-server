package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//net.Listen reserves the specified PORT for us
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// for { //used to keep server listening on the port infinitely
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	//Read and Write is a blocking connection, make sure to use it wisely
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	//localhost:4221/abcde returns 404 NOT FOUND,
	//localhost:4221/ returns 200 OK
	if strings.HasPrefix(string(buf), "GET / HTTP/1.1") {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
	conn.Close()
	// }
}
