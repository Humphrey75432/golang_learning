package main

import "fmt"

type MyInterface interface {
	Foo()
}

type MyStruct struct {
	intField int
	strField string
}

func describe(i MyInterface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func NewMyStruct() *MyStruct {
	return &MyStruct{1, "hello"}
}

func NewMyStruct2(arg int) *MyStruct {
	var obj = &MyStruct{intField: 1, strField: "hello"}
	if arg > 0 {
		return obj
	}
	return nil
}

func (s *MyStruct) Foo() {
	fmt.Println("Execute.%d,%s", s.intField, s.strField)
}

func main() {
	var i MyInterface
	describe(i)
	i.Foo()
}
