package main

import "fmt"

func main() {
	var s []string
	for i := 0; i < 9999; i++ {
		go func() {
			s = append(s, "Fuck you!")
		}()
	}
	// 写入Slice的次数被覆盖了，这是个有问题的代码，但是却不会报错
	fmt.Printf("Total fuck %d times.", len(s))
}
