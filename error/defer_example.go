package main

import "fmt"

func main() {
	fmt.Println("Main start")
	saveDivide(10, 0)
	saveDivide(10, 2)
	fmt.Println("Main end")
}

func saveDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Catch panic: %v\n", r)
			fmt.Println("Program has been recovered from panic, continue execute")
		}
	}()
	fmt.Printf("Trying %d / %d\n", a, b)
	if b == 0 {
		panic("Divisions by zero")
	}

	result := a / b
	fmt.Printf("Result: %d\n", result)
}
