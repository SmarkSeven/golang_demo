package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(2000 * time.Millisecond)
	go send(msg, done)

	for {
		select {
		case m, ok := <-msg:
			if ok {
				fmt.Println(m)
			}
		case <-until:
			fmt.Println("Stop")
			done <- true
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			// close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(100 * time.Millisecond)
		}
	}
}
