package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		// 使用atomic包可以更好的写出高并发代码
		atomic.AddUint64(&total2, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 开启2个goroutine
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()

	fmt.Println(total2)
}
