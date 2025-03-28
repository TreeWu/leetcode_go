package main

import "fmt"

/*
*
给定一个非负整数num。对于0 ≤ i ≤ num 范围中的每个数字i，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例2:

输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的__builtin_popcount）来执行此操作。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/counting-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(countBits(5))
}

/**
5  -> 0 ,1 ,2 ,3 ,4 ,5   ↓
      0,1,10,11,100,101  ↓
      [0,1,1,2,1,2]

*/
/**
对于 nums ,需要计算出所以小于 它 的 奇数 和 偶数 的比特位
在二进制分析 对于 5->101 它的高两位 与 2->10 一样，区别在于 最低位
每一个数 都可以表示为 f(x) = f(x/2) + x%2 = f(x>>1) + x&1
*/
func countBits(num int) []int {

	ans := make([]int, num+1)

	for i := 0; i <= num; i++ {
		ans[i] = ans[i>>1] + i&1
	}

	return ans
}
