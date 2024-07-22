package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

func run(conn net.Conn, buf []byte) {
	//to fetch out 'str' out of "/echo/{str}"
	time.Sleep(4 * time.Second)
	req := string(buf)

	sliceReq := strings.Split(req, " ")
	urlPath := sliceReq[1]
	if urlPath == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		return
	}
	if strings.HasPrefix(urlPath, "/echo/") {
		str := strings.Split(urlPath, "/")[2]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(str), str)))
		return
	}
	if strings.Contains(req, "/user-agent") {
		re := regexp.MustCompile(`User-Agent: (.*)\r\n`)
		matches := re.FindStringSubmatch(req)
		userAgentValue := matches[1]
		if len(matches) > 1 {
			// The captured group is in the second element of the slice
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(userAgentValue), userAgentValue)))
		}
		return
	}
	conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	conn.Close()
}
