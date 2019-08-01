package main

import "fmt"

/**
翻转链表
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

迭代法 ：翻转链表，我们可以肯定的是，不管怎么，链表的最后一个元素，就是 null
①：假如链表元素只有一个 n1->null，那么 既是  n1->null,
②：假如传入的是 n1->n2->null,那么我们先需要两个元素， pre{前元素}，cur{当前元素}， 将 cur.next = pre,就能完成最简单的翻转
	同时我们要 保存 下一个元素用做下次循环的 当前元素 ，
	对于下次循环来说，前元素既是本次循环的当前元素，而当前元素是本次当前元素的下一个
	一个是前变量，一个是当前变量，
在遍历列表时，将当前节点的 next 指针改为指向前一个元素。由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。在更改引用之前，还需要另一个指针来存储下一个节点。
*/
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil //前节点
	var cur = head          //当前节点

	for cur != nil { //判断当前节点是否是空，为空不继续
		tmp := cur.Next //保存当前节点的下一个节点，因为逆序的时候，当前节点的next要指向上一节点pre,所以要先保存。避免丢失
		cur.Next = pre  //当前节点的next指向pre，完成逆序，当第一次循环开始时，pre为null,
		pre = cur       // 将当前节点作为下次循环的 pre ,这样下次循环时的节点.next = pre
		cur = tmp       // 当前节点的next 作为下次 循环的 cur
	}

	return pre
}

/**
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

从结果来看，是将原链表分成两部分，同时逆转后半部分，然后前后两部分链表穿插
解题步骤：
1：找中间节点 （快慢指针法，一个走一步，一个走两步，快指针到链表尾的时候，慢指针就处于中间）
2：翻转后半部分
3：穿插
*/

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	tmp := head
	mid := FindMidNode(head)
	last := ReverseList(mid.Next)

	mid.Next = nil

	for tmp != nil && last != nil {
		t := tmp.Next
		tmp.Next = last
		last = last.Next
		tmp.Next.Next = t
		tmp = t
	}

	if last != nil {
		tmp.Next = last
	}
}

/**
找中间节点
快慢指针法，一个走一步，一个走两步，快指针到链表尾的时候，慢指针就处于中间
*/
func FindMidNode(head *ListNode) *ListNode {
	fast := head
	low := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
	}
	return low
}

func Println(head *ListNode) {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		Println(head.Next)
	}
}

func main() {
	size := 6
	head := &ListNode{0, nil}
	tmp := head
	for i := 1; i < size; i++ {
		tmp.Next = &ListNode{i, nil}
		tmp = tmp.Next
	}
	Println(head)
	ReorderList(head)
	fmt.Println()
	Println(head)
}
