package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func AddTowNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := 0
	result := &ListNode{0, nil}
	r := result
	for l1 != nil || l2 != nil {
		p, q, t := 0, 0, 0
		if l1 != nil {
			p = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			q = l2.Val
			l2 = l2.Next
		}

		t = p + q + tmp

		result.Val = t % 10
		tmp = t / 10

		if !(l1 == nil && l2 == nil) {
			result.Next = &ListNode{0, nil}
			result = result.Next
		}
	}

	if tmp != 0 {
		result.Next = &ListNode{tmp, nil}
	}
	return r
}
