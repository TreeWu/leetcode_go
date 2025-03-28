package main

/*
*
给定一个链表，删除链表的倒数第n个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		i++
	}
	if i == n {
		head = head.Next
	} else {

		i = i - n
		tmp = head
		for i > 1 {
			tmp = tmp.Next
			i--
		}
		tmp.Next = tmp.Next.Next
	}
	Println(head)
	return head
}

// 一趟扫描
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dmmy := &ListNode{0, head} // 建立一个辅助接点，处理当 n 等于链表长度的情况
	low, fast := dmmy, dmmy

	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		low = low.Next
	}
	low.Next = low.Next.Next
	return dmmy.Next
}

func main() {
	node := generateListNode(5)
	Println(node)
	removeNthFromEnd(node, 5)
}
