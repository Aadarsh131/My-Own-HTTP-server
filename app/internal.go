package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func checkBasic(conn net.Conn, buf []byte) {

	//localhost:4221/abcde returns 404 NOT FOUND,
	//localhost:4221/ returns 200 OK
	if strings.HasPrefix(string(buf), "GET / HTTP/1.1") {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}

//example request- GET /echo/abc HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n
//example response- HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nabc

/*
	Breakdown of response

// Status line
HTTP/1.1 200 OK
\r\n                          // CRLF that marks the end of the status line

// Headers
Content-Type: text/plain\r\n  // Header that specifies the format of the response body
Content-Length: 3\r\n         // Header that specifies the size of the response body, in bytes
\r\n                          // CRLF that marks the end of the headers

// Response body
abc                           // The string from the request
*/
func checkEcho(conn net.Conn, buf []byte) {
	// stringbuf := string(buf)
	// fmt.Println([]byte(" "))
	count := 0
	var tempSlice []string
	var tempString string
	for i := 0; i < len(buf); i++ {
		if buf[i] == 32 { //" "(space) in bytes is 32
			tempSlice = append(tempSlice, tempString)
			tempString = ""
			count++
			if count > 1 {
				break
			}
			continue
		}
		tempString += string(buf[i])
	}
	urlPath := tempSlice[1]

	matched, err := regexp.MatchString("/echo/", urlPath)
	if err != nil {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}
	if matched {
		res := strings.Split(urlPath, "/")[2]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(res), res)))
		return
	}
}
