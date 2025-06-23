package main

import (
	"fmt"
	"time"
)

func main() {
	finish := make(chan struct{})
	// 使用channel实现阻塞
	go func() {
		time.Sleep(time.Second)
		fmt.Println("over")
		finish <- struct{}{} // 将空结构体的示例传递给channel
	}()
	<-finish // 等待子线程结束,类似于Java中的t.join操作

	// 实现Set
	list := []int{1, 4, 3, 6, 5, 2, 3}
	set := make(map[int]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	if _, exist := set[3]; exist {
		fmt.Println("set contain element 3")
	}
}
