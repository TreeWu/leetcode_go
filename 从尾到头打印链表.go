package main

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

 

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
 

限制：

0 <= 链表长度 <= 10000

*/
func main() {

	type ListNode struct {
		Val  int
		Next *ListNode
	}

	type reversePrint func(head *ListNode) []int
	var v reversePrint
	v = func(head *ListNode) []int {
		if head != nil {
			return append(v(head.Next), head.Val)
		}
		return nil
	}
}
