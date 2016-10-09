package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "github.com:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
