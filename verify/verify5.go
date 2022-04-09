package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var cc = make(chan int)
	go client(cc)

	for {
		select {
		// Ensure channel has been graceful closed
		// Otherwise program will not exit when send kill signal to operator system
		case _, ok := <-cc:
			if ok == false {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("continue")
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
		}
	}
}

func client(c chan int) {
	defer close(c)
	for {
		err := processBusiness()
		if err != nil {
			c <- 0
			return
		}
		c <- 1
	}
}

func processBusiness() error {
	return errors.New("domo")
}
