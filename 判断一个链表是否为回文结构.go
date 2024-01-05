package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPail(ListNodeFromArray(1, 2, 3, 4, 5, 3, 2, 1)))
}
func isPail(head *ListNode) bool {
	// write code here
	one, two := head, head

	for two != nil {
		two = two.Next
		if two != nil && two.Next != nil {
			two = two.Next
			one = one.Next
		} else {
			break
		}
	}
	// 1 2 3 4 5
	// 链表翻转要两个指针，一个是翻转前的，一个是翻转后的，吧翻转前的第一个指针的next ,指向翻转后的第一个指针，
	//然后返回翻转后的第一个指针就可以了
	// 1 2 3 4
	// 2 1 3 4
	// 3 2 1 4
	// 4 3 2 1
	fanzhuang := func(head *ListNode) *ListNode {

		cur := head
		for cur.Next != nil {
			tmp := cur.Next.Next
			next := cur.Next
			next.Next = head
			cur.Next = tmp
			head = next
		}
		return head
	}

	one = fanzhuang(one.Next)

	for one != nil {

		if one.Val == head.Val {
			one = one.Next
			head = head.Next
		} else {
			return false
		}

	}
	return true
}
