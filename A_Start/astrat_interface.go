package astart

import (
	"container/heap"
)

type AStartNode interface {
	FindNeighbors() []AStartNode     //找到邻居节点
	SetParent(AStartNode)            //设置父节点
	Parent() AStartNode              //获取父节点
	G() float64                      //获取实际代价
	H() float64                      //获取启发代价
	SetG(float64)                    //设置实际代价
	SetH(float64)                    //设置启发代价
	Equal(AStartNode) bool           //判断两个节点是否相等
	NeighborCost(AStartNode) float64 //计算两个节点的代价
	Heuristic(AStartNode) float64    //计算启发代价
	F() float64                      //获取总代价
	PrintSolution()                  //打印路径
}

type DefaultNode struct {
	g, h, f float64
}

type AStartPriorityQueue []AStartNode

type AStartSolution struct {
	StartNode, TargetNode AStartNode
	OpenList              *AStartPriorityQueue
	InCloseList           func(AStartNode) bool
	PutToCloseList        func(AStartNode)
	OnNodeComputer        func(AStartNode)
}

func (d *DefaultNode) G() float64 {
	return d.g
}

func (d *DefaultNode) H() float64 {
	return d.h
}

func (d *DefaultNode) SetG(g float64) {
	d.g = g
}

func (d *DefaultNode) SetH(h float64) {
	d.h = h
}

func (d *DefaultNode) F() float64 {
	return d.g + d.h
}

func (A *AStartPriorityQueue) Len() int {
	return len(*A)
}

func (A *AStartPriorityQueue) Less(i, j int) bool {
	return (*A)[i].F() < (*A)[j].F()
}

func (A *AStartPriorityQueue) Swap(i, j int) {
	(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
}

func (A *AStartPriorityQueue) Push(x any) {
	*A = append(*A, x.(AStartNode))
}

func (A *AStartPriorityQueue) Pop() any {
	x := (*A)[len(*A)-1]
	*A = (*A)[:len(*A)-1]
	return x
}
func (A *AStartPriorityQueue) FindNode(node AStartNode) (AStartNode, int) {
	for index, n := range *A {
		if node.Equal(n) {
			return n, index
		}
	}
	return nil, 0
}

var _ heap.Interface = &AStartPriorityQueue{}

func (a AStartSolution) Solution() (AStartNode, bool) {
	if a.OpenList == nil {
		a.OpenList = &AStartPriorityQueue{}
	}
	if a.InCloseList == nil || a.PutToCloseList == nil {
		panic("InCloseList,PutToCloseList is nil")
	}
	heap.Init(a.OpenList)
	a.StartNode.SetH(a.StartNode.Heuristic(a.TargetNode))
	heap.Push(a.OpenList, a.StartNode)
	for a.OpenList.Len() != 0 {
		current := heap.Pop(a.OpenList).(AStartNode)
		if a.OnNodeComputer != nil {
			a.OnNodeComputer(current)
		}
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
