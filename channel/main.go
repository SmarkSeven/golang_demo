package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}
	go printCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
