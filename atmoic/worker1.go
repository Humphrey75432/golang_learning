package main

import (
	"fmt"
	"sync"
)

var total struct {
	// 借助加锁和解锁保证该语句在同一时刻只被一个线程访问
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		// 加锁，保证当前只有一个线程占有该资源
		total.Lock()
		total.value += i
		// 解锁，保证占有资源的线程释放后给其他线程机会抢占
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 线程1计算1 + 2 + ... + 100 = 5050
	go worker(&wg)
	// 线程2计算1 + 2 + ... + 100 = 5050
	go worker(&wg)
	wg.Wait()

	// 因此总值计算就为两个线程的和：10100
	fmt.Println(total.value)
}
