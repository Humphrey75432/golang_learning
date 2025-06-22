package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go doWork(ctx)
	time.Sleep(2 * time.Second)
	cancel()
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			fmt.Println("Doing some work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
