package main

import "fmt"

/**
插入排序
插入排序就是将数组分为有序区和无序区，在无序区中选择一个数字插入到有序区的合适位置
可以想象打牌的时候，从左到右整理牌时，就是在右边找一张牌，插入到左边去
*/
func InsertSort(arr []int) {
	// 可以想象，左边第一张牌，就是有序的，所以从1开始
	for i := 1; i < len(arr); i++ {
		insertValue := arr[i]
		j := i - 1 // 从 insertValue 左边第一位开始遍历
		for ; j >= 0; j-- {
			if arr[j] > insertValue { //查找小于insertValue的值，将大于它的值往右移，空出插入的位置
				arr[j+1] = arr[j]
			} else { //当值小于insertValue是，跳出循环，否则会再一次j--
				break
			}
		}
		arr[j+1] = insertValue
	}
}

/**
选择排序，类似冒泡排序，但是不是每次都交换元素，而是每次选出最小的值，进行交换
*/
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		for j := index; j < len(arr); j++ {
			if arr[j] < arr[index] { //循环遍历，找出最小值
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i] // 将最小值交换到左边
	}
}

/**
快排 分治法 双指针交换
*/
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := Index(arr, left, right)
	QuickSort(arr, left, index-1)
	QuickSort(arr, index+1, right)
}
func Index(arr []int, left, right int) int {
	pivot := arr[left]
	low, high := left, right
	for low != high {
		// 右指针左移，当指针指向的值比基准值小的时候，停止
		for low < high && arr[high] > pivot {
			high--
		}
		// 左指针右移，当指针指向的值比基准大的时候停止
		for low < high && arr[low] <= pivot {
			low++
		}
		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
		}
	}
	arr[low], arr[left] = arr[left], arr[low]
	return low
}

func main() {
	arr := []int{1, 3, 6, 8, 3, 2, 6, 3}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	var name string
	fmt.Scanln(&name)
}
