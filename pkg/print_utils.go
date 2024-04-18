package pkg

import (
	"container/list"
	"fmt"
)

// PrintSlice 打印切片
func PrintSlice[T any](nums []T) {
	fmt.Printf("%v", nums)
	fmt.Println()
}

// PrintList 打印列表
func PrintList(list *list.List) {
	if list.Len() == 0 {
		fmt.Print("[]\n")
		return
	}
	e := list.Front()
	// 强转为 string, 会影响效率
	fmt.Print("[")
	for e.Next() != nil {
		fmt.Print(e.Value, " ")
		e = e.Next()
	}
	fmt.Print(e.Value, "]\n")
}

// PrintMap 打印哈希表
func PrintMap[K comparable, V any](m map[K]V) {
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}
