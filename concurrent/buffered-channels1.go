package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	// ch1已经满了，3没办法放进去，会报错
	ch1 <- 3
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
