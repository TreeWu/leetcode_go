package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type Stack struct {
	Size int
	Head *ListNode
}

func (s *Stack) pop() (interface{}, error) {
	if s.Head == nil {
		return nil, errors.New("stack is empty")
	}
	tmp := s.Head
	s.Head = s.Head.Next
	s.Size = s.Size - 1
	return tmp.Val, nil
}
func (s *Stack) push(e interface{}) {
	node := &ListNode{e, s.Head}
	s.Size = s.Size + 1
	s.Head = node
}

func Println(head *ListNode) {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		Println(head.Next)
	}
	fmt.Println()
}
func (head *ListNode) Println() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print(" -> ")
		cur = cur.Next
	}
	fmt.Println()
}

func (s *Stack) peek() (interface{}, error) {
	if s.Head == nil {
		return nil, errors.New("stack is empty")
	}
	return s.Head.Val, nil
}

func generateListNode(size int) *ListNode {
	head := &ListNode{0, nil}
	tmp := head
	for i := 1; i < size; i++ {
		tmp.Next = &ListNode{i, nil}
		tmp = tmp.Next
	}
	return head
}

func ListNodeFromArray(array ...int) *ListNode {
	head := &ListNode{}
	pre := head
	for i := range array {
		pre.Next = &ListNode{Val: array[i]}
		pre = pre.Next
	}
	return head.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
