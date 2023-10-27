package main

import "fmt"

// 初始容量：选取一个合理的数组初始容量。在本示例中，我们选择 10 作为初始容量。
// 数量记录：声明一个变量 size ，用于记录列表当前元素数量，并随着元素插入和删除实时更新。根据此变量，我们可以定位列表尾部，以及判断是否需要扩容。
// 扩容机制：若插入元素时列表容量已满，则需要进行扩容。首先根据扩容倍数创建一个更大的数组，再将当前数组的所有元素依次移动至新数组。在本示例中，我们规定每次将数组扩容至之前的 2 倍。

/* 列表类简易实现 */
type myList struct {
	arrCapacity int   // 列表容量
	arr         []int // 数组（存储列表元素）
	arrSize     int   // 列表长度（即当前元素数量）
	extendRatio int   // 每次列表扩容的倍数
}

/* 构造函数 */
func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
}

/* 获取列表长度（即当前元素数量） */
func (l *myList) size() int {
	return l.arrSize
}

/*  获取列表容量 */
func (l *myList) capacity() int {
	return l.arrCapacity
}

/* 访问元素 */
func (l *myList) get(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	return l.arr[index]
}

/* 更新元素 */
func (l *myList) set(num, index int) {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	l.arr[index] = num
}

/* 尾部添加元素 */
func (l *myList) add(num int) {
	// 元素数量超出容量时，触发扩容机制
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = num
	// 更新元素数量
	l.arrSize++
}

/* 中间插入元素 */
func (l *myList) insert(num, index int) {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	// 元素数量超出容量时，触发扩容机制
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	// 将索引 index 以及之后的元素都向后移动一位
	for j := l.arrSize - 1; j >= index; j-- {
		l.arr[j+1] = l.arr[j]
	}
	l.arr[index] = num
	// 更新元素数量
	l.arrSize++
}

/* 删除元素 */
func (l *myList) remove(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	num := l.arr[index]
	// 索引 i 之后的元素都向前移动一位
	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	// 更新元素数量
	l.arrSize--
	// 返回被删除元素
	return num
}

/* 列表扩容 */
func (l *myList) extendCapacity() {
	// 新建一个长度为原数组 extendRatio 倍的新数组，并将原数组拷贝到新数组
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	// 更新列表容量
	l.arrCapacity = len(l.arr)
}

/* 返回有效长度的列表 */
func (l *myList) toArray() []int {
	// 仅转换有效长度范围内的列表元素
	return l.arr[:l.arrSize]
}

func main() {
	l := newMyList()

	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)

	fmt.Println(l.size())

	l.insert(2, 6)

}
