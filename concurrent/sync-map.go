package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {
	data := []string{"Fuck1", "Fuck2", "Fuck3", "Fuck4"}
	for i := 0; i < 4; i++ {
		go func(i int) {
			m.Store(i, data[i])
		}(i)
	}
	time.Sleep(time.Second)

	v, ok := m.Load(0)
	fmt.Printf("Load: %v, %v\n", v, ok)

	m.Delete(1)

	v, ok = m.LoadOrStore(1, "Fuck2")
	fmt.Printf("Load: %v, %v\n", v, ok)

	m.Range(func(key, value any) bool {
		fmt.Printf("Range: %v, %v\n", key, value)
		return true
	})
}
