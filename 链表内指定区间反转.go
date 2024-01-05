package main

import "fmt"

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
// 示例 2：
//
// 输入：head = [5], left = 1, right = 1
// 输出：[5]

func main() {

	for i := 1; i < -1; i++ {

		fmt.Println("aaa")
	}

	reverseBetween(ListNodeFromArray(3, 4), 1, 2).Println()
	reverseBetween(ListNodeFromArray(1, 2, 3, 4, 5), 2, 4).Println()

}
func reverseBetween(head *ListNode, m, n int) *ListNode {
	vmNode := &ListNode{Next: head}

	pre := vmNode

	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	// vm  1  3    2   4   5  6
	//pre     next

	next := pre.Next
	for ; m < n; m++ {
		tmp := pre.Next            // tmp 3 2 3 4 5
		pre.Next = next.Next       // 1 ->4 5 6  next = 3
		next.Next = next.Next.Next // 2.next = 5 6
		pre.Next.Next = tmp        // 1 4 3 2 5 6

	}
	return vmNode.Next
}
