package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

type MyMethod struct {
	Say  string
	Name string
}

func (my *MyMethod) SayHello() string {
	return "hello, " + my.Name
}

func (my *MyMethod) SayYouName(name string) string {
	return "goodbye, " + name
}

func main() {
	sweater := Inventory{"wool", 10}
	tmpl, err := template.New("test").Parse("{{ .Count }} of {{ .Material }}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweater)
	if err != nil {
		panic(err)
	}

	mine := &MyMethod{Say: "hello", Name: "student"}
	tmpl2, err := template.New("test").Parse("{{$str1 := .Say}}{{$str2 := .SayHello}}{{$str3 := .SayYouName .Name}}{{$str1}} {{$str2}}\n{{$str3}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl2.Execute(os.Stdout, mine)
	if err != nil {
		panic(err)
	}
}
