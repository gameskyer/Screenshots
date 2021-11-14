package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func main(){
	var l1 = new(ListNode)
	var l11 = new(ListNode)
	var l12 = new(ListNode)
	var l2 = new(ListNode)
	var l21 = new(ListNode)
	var l22 = new(ListNode)
	l1.Val = 8
	l11.Val = 9
	l12.Val = 7
	l12.Next = new(ListNode)
	l11.Next = l12
	l1.Next = l11
	l2.Val = 1
	l21.Val = 1
	l22.Val = 3
	l22.Next = new(ListNode)
	l21.Next = l22
	l2.Next = l21
	code := addTwoNumbers(l1,l2)
	fmt.Println(code)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headList := new(ListNode)
	head := headList
	num := 0

	for (l1 != nil || l2 != nil || num > 0) {
		head.Next =  new(ListNode)
		head = head.Next
		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		head.Val = (num) % 10
		num = num / 10
	}

	return headList.Next
}

