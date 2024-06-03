package greedy

import (
	"fmt"
	"testing"
)

func TestMaxCapacity(t *testing.T) {
	ht := []int{3, 8, 5, 2, 7, 7, 3, 4}

	// 貪婪演算法
	res := maxCapacity(ht)
	fmt.Println("最大容量為", res)
}
