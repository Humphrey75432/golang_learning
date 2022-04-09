package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var group sync.WaitGroup

	for i := 0; i < 5; i++ {
		group.Add(1)
		go func(j int) {
			defer group.Done()
			fmt.Printf("%-2d", j)
		}(i)
	}
	group.Wait()
}
