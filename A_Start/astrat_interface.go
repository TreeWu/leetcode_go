package main

import (
	"container/heap"
)

type AStartNodeInterface interface {
	FindNeighbors() []AStartNodeInterface
	SetParent(AStartNodeInterface)
	Parent() AStartNodeInterface
	G() float64
	H() float64
	SetG(float64)
	SetH(float64)
	Equal(AStartNodeInterface) bool
	NeighborCost(AStartNodeInterface) float64
	Heuristic(AStartNodeInterface) float64
	F() float64
	PrintSolution()
}

type DefaultNode struct {
	g, h, f float64
}

func (d DefaultNode) G() float64 {
	return d.g
}

func (d DefaultNode) H() float64 {
	return d.h
}

func (d DefaultNode) SetG(g float64) {
	d.g = g
}

func (d DefaultNode) SetH(h float64) {
	d.h = h
}

func (d DefaultNode) F() float64 {
	return d.g + d.h
}

type AStartPriorityQueue []AStartNodeInterface

func (A AStartPriorityQueue) Len() int {
	return len(A)
}

func (A AStartPriorityQueue) Less(i, j int) bool {
	return A[i].G()+A[i].H() < A[j].G()+A[j].H()
}

func (A AStartPriorityQueue) Swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}

func (A *AStartPriorityQueue) Push(x any) {
	*A = append(*A, x.(AStartNodeInterface))
}

func (A *AStartPriorityQueue) Pop() any {
	x := (*A)[len(*A)-1]
	*A = (*A)[:len(*A)-1]
	return x
}
func (A *AStartPriorityQueue) FindNode(node AStartNodeInterface) (AStartNodeInterface, int) {
	for index, n := range *A {
		if node.Equal(n) {
			return n, index
		}
	}
	return nil, 0
}

var _ heap.Interface = &AStartPriorityQueue{}

type AStartSolution struct {
	StartNode, TargetNode AStartNodeInterface
	OpenList              *AStartPriorityQueue
	InCloseList           func(AStartNodeInterface) bool
	PutToCloseList        func(AStartNodeInterface)
}

func (a AStartSolution) Solution() (AStartNodeInterface, bool) {
	if a.OpenList == nil {
		a.OpenList = &AStartPriorityQueue{}
	}
	if a.InCloseList == nil || a.PutToCloseList == nil {
		panic("InCloseList,PutToCloseList is nil")
	}
	heap.Init(a.OpenList)
	heap.Push(a.OpenList, a.StartNode)
	for a.OpenList.Len() != 0 {
		current := heap.Pop(a.OpenList).(AStartNodeInterface)
		if current.Equal(a.TargetNode) {
			return current, true
		}
		if a.InCloseList(current) {
			continue
		}
		a.PutToCloseList(current)
		neighbors := current.FindNeighbors()
		for i := range neighbors {
			neighbor := neighbors[i]
			if a.InCloseList(neighbor) {
				continue
			}
			g := current.G() + current.NeighborCost(neighbor)
			node, index := a.OpenList.FindNode(neighbor)
			if node == nil {
				neighbor.SetG(g)
				neighbor.SetH(current.Heuristic(a.TargetNode))
				neighbor.SetParent(current)
				heap.Push(a.OpenList, neighbor)
			} else if node.G() > g {
				node.SetG(g)
				node.SetParent(current)
				heap.Fix(a.OpenList, index)
			}

		}

	}
	return nil, false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
