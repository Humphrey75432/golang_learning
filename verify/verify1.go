package main

import (
	"fmt"
	"time"
)

func DoWaiter(name string, second int, signal chan int) {
	for {
		select {
		case <-time.Tick(time.Duration(second) * time.Second):
			fmt.Println(name, " is ready!")
		case <-signal:
			fmt.Println(name, " close goroutine.")
			return
		}
	}
}

func main() {
	var signal1 chan int = make(chan int)
	var signal2 chan int = make(chan int)

	defer close(signal1)
	defer close(signal2)

	go DoWaiter("Tea", 2, signal1)
	go DoWaiter("Coffee", 1, signal2)

	fmt.Println("main() is waiting...")
	time.Sleep(4 * time.Second)

	signal1 <- 1
	signal2 <- 1
	time.Sleep(1 * time.Second)
}
