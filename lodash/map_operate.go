package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	//	Keys
	keys := lo.Keys(map[string]int{"foo": 1, "bar": 2})
	fmt.Printf("Keys:%v\n", keys)

	//	UniqKeys
	uniqKeys := lo.UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 2})
	fmt.Printf("UniqKeys:%v\n", uniqKeys)

	//	HasKey
	exists := lo.HasKey(map[string]int{"foo": 1, "bar": 2}, "bar")
	fmt.Printf("HasKey:%t\n", exists)
	exists = lo.HasKey(map[string]int{"foo": 1, "bar": 2}, "baz")
	fmt.Printf("HasKey:%t\n", exists)

	//	Values
	values := lo.Values(map[string]int{"foo": 1, "bar": 2})
	fmt.Printf("Values:%v\n", values)

	//	UniqValues
	uniqValues := lo.UniqValues(map[string]int{"foo": 1, "bar": 2, "baz": 1})
	fmt.Printf("UniqValues:%v\n", uniqValues)

	// ValueOr
	value := lo.ValueOr(map[string]int{"foo": 1, "bar": 2}, "foo", 32)
	fmt.Printf("ValueOr:%d\n", value)
	value = lo.ValueOr(map[string]int{"foo": 1, "bar": 2}, "fuck", 32)
	fmt.Printf("ValueOr:%d\n", value)

	//	PickBy
	picked := lo.PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})
	fmt.Printf("PickBy:%v\n", picked)

	//	PickByKeys
	pickedByKeys := lo.PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	fmt.Printf("PickByKeys:%v\n", pickedByKeys)
}
