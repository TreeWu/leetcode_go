package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
179. 最大数
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。



示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
示例 3：

输入：nums = [1]
输出："1"
示例 4：

输入：nums = [10]
输出："10"


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109
 */

func main() {
fmt.Println(largestNumber([]int{3,30,34,5,9}))
}
func largestNumber(nums []int) string {
	ans:=""
	var ss []string
	for i := range nums {
		itoa := strconv.Itoa(nums[i])
		ss=append(ss,itoa)
	}
	sort.Slice()
	sort.Strings(ss)


	for i := range ss {
		ans+=ss[len(ss) -i-1]
	}

	return  ans
}
