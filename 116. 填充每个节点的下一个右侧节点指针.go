package main

import (
	"container/list"
	"fmt"
)

/**
116. 填充每个节点的下一个右侧节点指针
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



示例：



输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}

输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}

解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。


提示：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/
func main() {

	type Node struct {
		Val   int
		Left  *Node
		Right *Node
		Next  *Node
	}

	var connect func(root *Node) *Node
	connect = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		l := list.New()
		l.PushBack(root)
		nextLen, curLen := 0, 1
		var pre *Node
		for l.Len() != 0 {
			front := l.Front()
			l.Remove(front)
			node := front.Value.(*Node)
			curLen--
			if pre == nil {
				pre = node
			} else {
				pre.Next = node
				pre = node
			}
			if curLen == 0 {
				pre = nil
			}
			if node.Left != nil {
				l.PushBack(node.Left)
				nextLen++
			}
			if node.Right != nil {
				l.PushBack(node.Right)
				nextLen++
			}
			if curLen == 0 {
				curLen = nextLen
				nextLen = 0
			}
		}
		return root
	}

	connect = func(root *Node) *Node {

		if root == nil {
			return nil
		}
		l := list.New()
		l.PushBack(root)
		nextLen, curLen := 0, 1
		var pre *Node
		for l.Len() != 0 {
			front := l.Front()
			l.Remove(front)
			node := front.Value.(*Node)
			curLen--
			if pre == nil {
				pre = node
			} else {
				pre.Next = node
				pre = node
			}
			if curLen == 0 {
				pre = nil
			}
			if node.Left != nil {
				l.PushBack(node.Left)
				nextLen++
			}
			if node.Right != nil {
				l.PushBack(node.Right)
				nextLen++
			}
			if curLen == 0 {
				curLen = nextLen
				nextLen = 0
			}
		}
		return root

	}
	node := connect(&Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}})
	fmt.Println(node)
}
