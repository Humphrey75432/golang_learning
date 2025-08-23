package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// Substring
	sub := lo.Substring("hello", -4, 3)
	fmt.Printf("Substring:%s\n", sub)

	// ChunkString
	chunkString := lo.ChunkString("123456", 2)
	fmt.Printf("ChunkString:%v\n", chunkString)

	// PascalCase
	str := lo.PascalCase("hello_world")
	fmt.Printf("PascalCase:%s\n", str)

	// CamelCase
	camelCase := lo.CamelCase("hello_world")
	fmt.Printf("CamelCase:%s\n", camelCase)

	// KebabCase
	kebabCase := lo.KebabCase("helloWorld")
	fmt.Printf("KebabCase:%s\n", kebabCase)

	// SnakeCase
	snakeCase := lo.SnakeCase("HelloWorld")
	fmt.Printf("SnakeCase:%s\n", snakeCase)

	ch := make(chan int, 42)
	for i := 0; i <= 10; i++ {
		ch <- i
	}

	// ChannelDispatcher: Distribute message from input channels from N child channel
	// Close events are propagated to children
	children := lo.ChannelDispatcher(ch, 5, 10, lo.DispatchingStrategyRoundRobin[int])

	consumer := func(c <-chan int) {
		for {
			msg, ok := <-c
			if !ok {
				println("channel closed")
				break
			}
			println(msg)
		}
	}

	for i := range children {
		go consumer(children[i])
	}
}
