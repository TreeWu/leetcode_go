package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 入堆
func (h *MaxHeap) Push(x int) {
	h.arr = append(h.arr, x)
	h.up(len(h.arr) - 1)
}

// 这是用于插入元素到堆中的操作。
// 当你向堆中插入一个元素时，它被添加到堆的末尾，然后通过向上堆化的过程，将它移到合适的位置，以满足最大堆性质。
// 这通常涉及与父节点比较并交换的过程，直到堆性质被恢复。

// 如果堆本来就是有序的，那么只需要将插入的元素，和父节点比较，然后交换就能保证堆的排序，类似冒泡一样
func (h *MaxHeap) up(index int) {
	parent := (index - 1) / 2
	if parent >= 0 && h.arr[index] > h.arr[parent] {
		h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
		h.up(parent)
	}
}

func (h *MaxHeap) Size() int {
	return len(h.arr)
}
func (h *MaxHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MaxHeap) PopMax() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}
	max := h.arr[0]

	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	if !h.IsEmpty() {
		h.arr[0] = last
		h.down(0)
	}

	return max, true
}

func (h *MaxHeap) down(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	largest := index

	if leftChild < h.Size() && h.arr[leftChild] > h.arr[largest] {
		largest = leftChild
	}

	if rightChild < h.Size() && h.arr[rightChild] > h.arr[largest] {
		largest = rightChild
	}

	if largest != index {
		h.arr[index], h.arr[largest] = h.arr[largest], h.arr[index]
		h.down(largest)
	}
}

func heapSort(arrs []int) []int {

	heap := NewMaxHeap()

	// 插入元素到最大堆
	for _, el := range arrs {
		heap.Push(el)
	}
	max, ok := heap.PopMax()
	i := 0
	for ok {
		max, ok = heap.PopMax()
		arrs[i] = max
		i++
	}
	return arrs
}

func main() {
	heap := NewMaxHeap()
	elements := []int{4, 10, 3, 5, 1}

	fmt.Printf("排序后: %v \n", heapSort(elements))

	// 插入元素到最大堆
	for _, el := range elements {
		heap.Push(el)
	}

	// 输出堆中的最大值
	max, ok := heap.PopMax()
	if ok {
		fmt.Printf("最大值: %d\n", max)
	}

	// 输出剩余元素
	fmt.Printf("剩余元素: %v\n", heap.arr)
}
