package main

import (
	"fmt"
)

// 保留最大元素的 topK 堆
// 小堆顶
type TopKHeap struct {
	arr  []int
	size int
}

func NewTopKHeap(size int) *TopKHeap {
	return &TopKHeap{
		arr:  make([]int, size),
		size: 0,
	}
}

func (t *TopKHeap) Push(val int) {
	// 当堆未满时，需要插入元素到尾部，并且重新向上堆化
	if t.size < len(t.arr) {
		t.arr[t.size] = val
		t.up(t.size)
		t.size++

		// 如果插入的元素比 最小值大，替换堆顶并向下堆化
	} else if t.arr[0] < val {
		t.arr[0] = val
		t.down(0)
	}
}

func (t *TopKHeap) up(index int) {
	parent := (index - 1) / 2

	for parent >= 0 && t.arr[parent] > t.arr[index] {
		t.arr[parent], t.arr[index] = t.arr[index], t.arr[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (t *TopKHeap) down(index int) {
	left := index*2 + 1
	right := index*2 + 2
	parent := index

	if left < t.size && t.arr[left] < t.arr[parent] {
		parent = left
	}

	if right < t.size && t.arr[right] < t.arr[parent] {
		parent = right
	}

	if index != parent {
		t.arr[index], t.arr[parent] = t.arr[parent], t.arr[index]
		t.down(parent)
	}
}

func main() {
	heap := NewTopKHeap(5)
	arrs := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := range arrs {
		heap.Push(arrs[i])
	}

	fmt.Println(heap.arr)
}

//            0
//    1              2
// 3     8       9       7
//6
