package main

import "fmt"

func main() {
	/* 初始化栈 */
	// 在 Go 中，推荐将 Slice 当作栈来使用
	var stack []int

	/* 元素入栈 */
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 4)
	stack = append(stack, 5)

	/* 访问栈顶元素 */
	peek := stack[len(stack)-1]

	fmt.Println("栈顶元素为：", peek) // 5

	/* 元素出栈 */
	stack = stack[:len(stack)-1]
	fmt.Println("出栈后栈为：", stack) // [1 3 2 4]

	/* 获取栈的长度 */
	size := len(stack)
	fmt.Println("栈的长度为：", size) // 4

	/* 判断是否为空 */
	isEmpty := len(stack) == 0
	fmt.Println("栈是否为空：", isEmpty) // false

}
