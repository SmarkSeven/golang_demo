package main

import (
	"fmt"
	"log"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		// 有以下几种情况：
		// 1.没有数据
		// 2.有部分数据
		// 3.有足够多的数据
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		//  c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		}
		// start a new goroutine to handle the new connection
		go handleConn(c)
	}
}
