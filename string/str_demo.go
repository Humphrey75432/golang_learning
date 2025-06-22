package main

import "fmt"

func calcTotalPage(count int64) int {
	totalPage := count / 20
	if count%20 == 0 {
		return int(totalPage)
	}
	return int(totalPage + 1)
}

func main() {
	fmt.Println(calcTotalPage(4112))
}
