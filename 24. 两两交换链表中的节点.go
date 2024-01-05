package main

import (
	"fmt"
)

/*
*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func main() {

	type ListNode struct {
		Val  int
		Next *ListNode
	}

	print := func(l *ListNode) {
		var list []int
		for l != nil {
			list = append(list, l.Val)
			l = l.Next
		}
		fmt.Println(list)

	}

	swapPairs := func(head *ListNode) *ListNode {
		h := head
		pre := &ListNode{Next: head}
		for pre.Next != nil {
			if pre.Next == nil || pre.Next.Next == nil {
				break
			}
			if h == head {
				h = head.Next
			}
			next := pre.Next.Next
			pre.Next.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
			pre = next.Next
		}
		return h
	}

	print(swapPairs(&ListNode{1, &ListNode{Val: 2, Next: &ListNode{3, &ListNode{Val: 4}}}}))

}
