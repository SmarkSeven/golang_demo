package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := &count{}
	var file string
	for _, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			println(filename)
			if err := compress(filename); err == nil {
				counter.Increment()
			}
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Printf("Compressed %d files\n", counter.Sum())
}

type count struct {
	sync.Mutex
	num int
}

func (c *count) Increment() {
	c.Lock()
	c.num++
	c.Unlock()
}
func (c *count) Sum() int {
	return c.num
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
