package main

/**
303. 区域和检索 - 数组不可变
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。
*/
func main() {

}

type NumArray struct {
	sumMap []int
}

func Constructor(nums []int) NumArray {
	array := NumArray{
		sumMap: make([]int, len(nums)),
	}

	for i := range nums {
		if i == 0 {
			array.sumMap[i] = nums[i]
			continue
		}
		array.sumMap[i] = array.sumMap[i-1] + nums[i]
	}
	return array
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sumMap[j]
	}
	return this.sumMap[j] - this.sumMap[i-1]

}
