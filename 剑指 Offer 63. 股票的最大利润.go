package main

import "fmt"

/**
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？



示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


限制：

0 <= 数组长度 <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
解题思路  双指针

因为 利润一定是 后面的 减前面的，所以可以定义两个指针，left ,right ，right-left 即为 利润
同时我们应该灵活的 改变 left 的位置 ，什么时候改变呢？
假设 有 4 2 1  3  6  ,  显然 6 - 1 是最大的，在 right 往右移动的过程中，如果 right值比 left值小，那么就将left = right
因为 6 减 1 前面所有的数，但是比5 小的
*/
func main() {
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
}
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	left, right := 0, 1
	result := 0
	for right < len(prices) {
		if prices[right] < prices[left] {
			left = right
		} else {
			if prices[right]-prices[left] > result {
				result = prices[right] - prices[left]
			}
		}
		right++
	}
	return result
}

/**
动态规划， 到 i 日的最大利润，等于  i-1 日的最大利润 和 当前股票价值减去之前最低价格  的最大值
即  dp[i] = max（ dp[i-1] ， prices[i]-min(prices[:i]))
*/

func maxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	ans := 0
	min := prices[0]
	for index := 0; index < len(prices); index++ {
		if prices[index] < min {
			min = prices[index]
		}
		ans = max(ans, prices[index]-min)
	}

	return ans
}
