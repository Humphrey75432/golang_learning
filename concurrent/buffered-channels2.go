package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	// 输出1，剩余一个空位置
	fmt.Println(<-ch1)
	ch1 <- 3
	// 输出2，此时3也在通道中
	fmt.Println(<-ch1)
	// 输出3，此时通道为空
	fmt.Println(<-ch1)
}
