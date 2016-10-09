package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	// conn, err := net.DialTimeout("tcp", ":8888", time.Second)
	if err != nil {
		fmt.Println("dail error:", err)
		return
	}
	defer conn.Close()
	// read or write on conn
}
func client2() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client2.go YOUR_CONTENT")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	data := os.Args[1]
	conn.Write([]byte(data))

	time.Sleep(time.Second * 10000)
}
