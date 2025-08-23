package main

import (
	"fmt"
	"github.com/samber/lo"
	lom "github.com/samber/lo/mutable"
	"strconv"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	//	Filter
	even := lo.Filter([]int{1, 2, 3, 4}, func(item int, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("Even:%v\n", even)

	odd := lo.Filter([]int{1, 2, 3, 4}, func(item int, index int) bool {
		return item%2 != 0
	})
	fmt.Printf("Odd:%v\n", odd)

	even = lom.Filter([]int{1, 2, 3, 4}, func(item int) bool {
		return item%2 == 0
	})
	fmt.Printf("Even:%v\n", even)

	// Map
	res := lo.Map([]int64{1, 2, 3, 4}, func(item int64, index int) string {
		return fmt.Sprintf("toString(%d)", item)
	})
	fmt.Printf("Map result:%s\n", res)

	list := []int{1, 2, 3, 4}
	lom.Map(list, func(item int) int {
		return item * 2
	})
	fmt.Printf("List:%v\n", list)

	//	UniqMap
	users := []User{
		{Name: "Alex", Age: 10},
		{Name: "Alex", Age: 12},
		{Name: "Bob", Age: 11},
		{Name: "Alice", Age: 20},
	}
	names := lo.UniqMap(users, func(u User, index int) string {
		return u.Name
	})
	fmt.Printf("UniqMap Names:%v\n", names)

	//	FilterMap
	matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(item string, _ int) (string, bool) {
		if strings.HasSuffix(item, "pu") {
			return "xpu", true
		}
		return "", false
	})
	fmt.Printf("FilterMap:%v\n", matching)

	// FlatMap
	flatMaps := lo.FlatMap([]int64{0, 1, 2}, func(item int64, _ int) []string {
		return []string{
			strconv.FormatInt(item, 10),
			strconv.FormatInt(item, 10),
		}
	})
	fmt.Printf("FlatMap:%v\n", flatMaps)

	// Reduce
	sig := lo.Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	fmt.Printf("Reduce:%d\n", sig)

	//	ForEach
	lo.ForEach([]string{"hello", "world", "mother"}, func(item string, _ int) {
		println(item)
	})

	// Times
	times := lo.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Printf("Times:%v\n", times)

	//	Uniq
	uniqValues := lo.Uniq([]int{1, 2, 2, 1})
	fmt.Printf("Uniq:%v\n", uniqValues)

	//	Chunk
	chunks := lo.Chunk([]int{1, 2, 3, 4, 5, 6}, 3)
	fmt.Printf("Chunk:%v\n", chunks)

	//	PartitionBy
	partitions := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Printf("PartitionBy:%v\n", partitions)

	//	Flatten
	flat := lo.Flatten([][]int{{0, 1}, {2, 3, 4, 5}})
	fmt.Printf("Flatten:%v\n", flat)

	//	InterLeave
	interleaved := lo.Interleave([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
	fmt.Printf("Interleaved:%v\n", interleaved)

	// Shuffle
	shuffleList := []int{1, 2, 3, 4, 5, 6}
	lom.Shuffle(shuffleList)
	fmt.Printf("Shuffle:%v\n", shuffleList)

	//	Reverse
	orderList := []int{1, 2, 3, 4, 5, 6}
	lom.Reverse(orderList)
	fmt.Printf("Reverse:%v\n", orderList)

}
