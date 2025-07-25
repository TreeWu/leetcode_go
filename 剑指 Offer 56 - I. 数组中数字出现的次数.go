package main

import "fmt"

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。



示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
回忆简单版本的找 唯一不同的数字，可以直接使用异或运算，因为 异或 是相同为0，不同为 1 ，所以对所有数字一起异或，那么剩下来的数，就是 唯一一个不同的 数

这道题需要找 不同的两个数字 ，如果把所有的数字都进行异或，那么得到的只能是 这两个数字的 异或和
如果我们能想到办法，可以将数组分组，而且相同的数字分到同一局，不同的数字不同组，那么单独对两个组的数据异或，那么就能得出 两个不同的数了
那么怎么分组呢？？？
回想奇数偶数分组，因为偶数的最低位是 0 ，所以 n&1 ==1 那么就是奇数，如果 n&1 == 0 那么就是偶数 这里的 & 后面的 1 ，就是我们用来划分奇偶性的 关键位
回到我们的解题中，我们需要找到吧数组分组的 关键位，这个 关键位 需要同时满足两个条件
 1、能把不同的数字分在不同的组
 2、把相同的数字分在一样的组
异或的规则是相同为 0 ，不同为 1，所以我们可以在 全数组的异或和 中， 找出 1 的最低位，作为 关键位
将所有的数字再与改 关键位 异或 ，等于0的一组，其他另一组，最终就能求出 不同的另位数
假设有 数字  1    3     1      4
         0001   0011  0001   0100
异或和为 0111 从异或和中选择 最低位 0001 作为 关键位
再与上面 的所有数字异或，显然 1,3,1 分组在一组 求出 3， 4单独一组，求出4
所以 3,4 答案
*/
func singleNumbers(nums []int) []int {
	result := []int{0, 0}

	var sum int
	// 先求总数组的异或和
	for num := range nums {
		sum ^= nums[num]
	}
	// 找出 关键位
	//例如 在 0110 找一个 为 1 的位，先用 0001 与运算 ，得出 0000，为0，继续
	//把 0001 右移得到 0010，再与运算，得出 0010，不为0，所以 0010就是关键位
	keyPoint := 1
	for keyPoint&sum == 0 {
		keyPoint = keyPoint << 1
	}

	for index := range nums {
		num := nums[index]
		if num&keyPoint == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}
	return result
}

func main() {
	fmt.Println(singleNumbers([]int{1, 3, 1, 4}))
}
