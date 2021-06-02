package main

/**
剑指 Offer 09. 用两个栈实现队列
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
通过次数151,868提交次数209,203
*/
func main() {

}

type CQueue struct {
	stock1 []int
	stock2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stock2 = append(this.stock2, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stock1) == 0 && len(this.stock2) == 0 {
		return -1
	}
	if len(this.stock1) == 0 {
		this.stock1, this.stock2 = this.stock2, this.stock1
	}
	i := this.stock1[0]
	this.stock1 = this.stock1[1:len(this.stock1)]
	return i
}
