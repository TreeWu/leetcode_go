package main

func main() {

}

/*
* FindMedianSortedArrays 给定两个大小为 m 和 n 的有序数组nums1 和nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为O(log(m + n))。

你可以假设nums1和nums2不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

中位数: 将一个有序的数组，分为左右长度相等的数
如 1,2,3 中位数为 2， 1,2,3,4  中位数为 （2+3）/2 = 2.5
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	start, end, mid := 0, m, (m+n+1)/2
	for start <= end {
		i := (start + end) / 2
		j := mid - 1
		if i < end && nums2[j-1] > nums1[i] {
			// i偏小，需要右移
			start = i + 1
		} else if i > start {

		}
	}
	return result
}
