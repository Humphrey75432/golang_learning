package main

import (
	"fmt"
	"runtime"
	"time"
)

type Data struct {
	d [1024 * 100]byte
	o *Data
}

func foo() {
	var a, b Data
	a.o = &b
	b.o = &a
	runtime.SetFinalizer(&a, func(d *Data) {
		fmt.Printf("a %p final.\n", d)
	})
	runtime.SetFinalizer(&b, func(d *Data) {
		fmt.Printf("b %p final.\n", d)
	})
}

func main() {
	for {
		foo()
		time.Sleep(time.Millisecond)
	}
}
