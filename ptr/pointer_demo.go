package main

import "fmt"

func getPointerValue() *int {
	value := 42
	return &value
}

func main() {
	ptr := getPointerValue()
	fmt.Println("Pointer value: ", *ptr)
	fmt.Println("Pointer address: ", ptr)
}
