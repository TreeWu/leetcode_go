package main

import "fmt"

/**
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/
func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println(arr)
	moveZeroes(arr)
	fmt.Println(arr)

}

func moveZeroes(nums []int) {

	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroIndex == -1 {
				zeroIndex = i
			}
		} else {
			if zeroIndex != -1 {
				nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
				for j:=zeroIndex;j<=i;j++ {
					if nums[j]==0 {
						zeroIndex = j
						break
					}
				}
			}
		}
	}

}
