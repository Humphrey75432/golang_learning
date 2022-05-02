package main

import "fmt"

func main() {
	s := make(map[string]string)

	for i := 0; i < 99; i++ {
		go func() {
			s["Fuck"] = "You"
		}()
	}
	fmt.Printf("Totally fuck %d times", len(s))
}
