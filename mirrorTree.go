package main

import (
	"fmt"
)

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/

/**
解题思路 一个对称的二叉树，从根节点开始，它的左边，和右边应该就是对称的
并且，它自己和自己也是对称的
*/

type Queue struct {
	arrary []interface{}
	size   int
}

func (q *Queue) push(v interface{}) {
	q.arrary = append(q.arrary, v)
	q.size++
}
func (q *Queue) len() int {
	return q.size
}
func (q *Queue) poll() interface{} {
	if q.len() == 0 {
		return nil
	}
	v := q.arrary[q.size-1]
	q.arrary = q.arrary[:q.size-1]
	q.size--
	return v
}

func main() {

	i := []int{1, 2, 3, 4}

	len := 4

	fmt.Println(i[len-1])

	i = i[:len-1]

	fmt.Println(i)

}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		if t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left) {
			return true
		}
	}
	return false
}

/**
第二种解法
把要对比的对称节点，放入队列中，然后两两取出，判断节点值是否相等

*/

func isMirrorArray(root *TreeNode) bool {
	q := Queue{}
	q.push(root)
	q.push(root)

	for q.len() != 0 {
		t1 := q.poll().(*TreeNode)
		t2 := q.poll().(*TreeNode)

		if t1 == nil && t2 == nil { //都为空
			continue
		}
		if t1 != nil && t2 != nil { //都不为空
			if t1.Val != t2.Val { // 值相等
				return false
			}
			q.push(t1.Left)
			q.push(t2.Right)
			q.push(t1.Right)
			q.push(t2.Left)
			continue //下一轮遍历
		}
		return false //一个为空，一个不为空
	}
	return true

}
