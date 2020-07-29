package main

func trap(height []int) int {
	minIndex, maxIndex, sum := 0, 0, 0
	for index, value := range height {
		index := index
		value := value

		if height[index] > height[maxIndex] {
			maxIndex = index
		}
		if height[index] < height[minIndex] {
			minIndex = index
		}
	}
	return 0
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
