package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Value: val,
		Next:  nil,
	}
}

/* 在链表的节点 l2 之后插入节点 P */
func insertNode(l2 *ListNode, p *ListNode) {
	n := l2.Next
	p.Next = n
	l2.Next = p
}

/* 删除链表的节点 l2 之后的首个节点 */
func removeItem(l2 *ListNode) {
	if l2.Next == nil {
		return
	}
	l2.Next = l2.Next.Next
}

/* 访问元素，需要从头节点出发，逐个向后遍历，直至找到目标节点 */
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head.Next == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

/* 在链表中查找值为 target 的首个节点 */
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Value == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

func main() {
	l0 := NewListNode(0)
	l1 := NewListNode(1)
	l2 := NewListNode(2)
	l3 := NewListNode(3)
	l4 := NewListNode(4)
	l5 := NewListNode(5)

	l0.Next = l1
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	fmt.Println(l0)

	insertNode(l2, l5)
	fmt.Println(l0)

	removeItem(l2)
	fmt.Println(l0)

	node := access(l0, 3)
	fmt.Println(node.Value)
}
