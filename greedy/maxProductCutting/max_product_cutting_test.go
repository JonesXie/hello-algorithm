package greedy

import (
	"fmt"
	"testing"
)

func TestMaxProductCutting(t *testing.T) {
	n := 58
	// 貪婪演算法
	res := maxProductCutting(n)
	fmt.Println("最大切分乘積為", res)
}
