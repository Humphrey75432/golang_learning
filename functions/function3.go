package main

import "fmt"

func fibonacci() func() int {
	// 定义初始值
	back1, back2 := 0, 1
	return func() int {
		// 在函数闭包中实现斐波那契数列的算法
		temp := back1
		// 后一个数等于前两个数字的相加
		back1, back2 = back2, back1+back2
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
