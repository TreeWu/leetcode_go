package main

import "github.com/pkg/errors"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type Stack struct {
	Size int
	Head *ListNode
}

func (s *Stack) pop() (interface{},error) {
	if s.Head ==nil {
		return nil,errors.New("stack is empty")
	}
	tmp := s.Head
	s.Head = s.Head.Next
	s.Size = s.Size - 1
	return tmp.Val,nil
}
func (s *Stack) push(e interface{}) {
	node := &ListNode{e, s.Head}
	s.Size = s.Size + 1
	s.Head = node
}

func (s *Stack) peek() (interface{},error) {
	if s.Head ==nil {
		return nil,errors.New("stack is empty")
	}
	return s.Head.Val,nil
}