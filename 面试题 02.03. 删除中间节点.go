package main

/**
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。

 

示例：

输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-middle-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

// 看是否有下一节点，没有的话就把 值置空
// 如果有的话，那就把节点的值用next的值覆盖，next 用 next.next 覆盖
// 把 我 变成 你 ，我就不存在了
func deleteNode(node *ListNode) {
	if node.Next == nil {
		node.Val = nil
		node.Next = nil
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
