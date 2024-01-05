package astart

import (
	"container/heap"
	"log"
)

type Node struct {
	X, Y    int
	F, G, H int
	Parent  *Node
}

// 优先队列，用于存放待检查的节点
// 实际就是小顶堆
type PriorityQueue []*Node

var _ heap.Interface = &PriorityQueue{}

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].F < p[j].F
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(*Node))
}

func (p *PriorityQueue) NodeInQueueu(node *Node) (*Node, int) {
	for index, n := range *p {
		if n.X == node.X && n.Y == node.Y {
			return n, index
		}
	}
	return nil, 0
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

// 启发函数，预估从当前节点到目标节点的距离
type Heuristic func(int, int, int, int) int

// 计算从当前节点到邻居节点的代价
type NeighborCost func(int, int, int, int) int

// 查找邻居节点
type FindNeighbor func(int, int) []*Node
type Solution struct {
	StartX, StartY   int
	TargetX, TargetY int
	OpenList         *PriorityQueue
	CloseList        map[int]map[int]bool
	// 启发函数，预估从当前节点到目标节点的距离
	Heuristic    Heuristic
	FindNeighbor FindNeighbor
	// 计算从当前节点到邻居节点的代价
	NeighborCost NeighborCost
}

func (a Solution) Solution() (*Node, bool) {

	// 初始化起始点
	if a.OpenList == nil {
		a.OpenList = &PriorityQueue{}
	}
	a.OpenList.Push(&Node{
		X:      a.StartY,
		Y:      a.StartX,
		F:      a.Heuristic(a.StartX, a.StartY, a.TargetX, a.TargetY),
		G:      0,
		H:      a.Heuristic(a.StartX, a.StartY, a.TargetX, a.TargetY),
		Parent: nil,
	})
	heap.Init(a.OpenList)
	for a.OpenList.Len() != 0 {
		current := heap.Pop(a.OpenList).(*Node)
		log.Printf("当前节点 %d,%d \n", current.X, current.Y)
		// 如果当前节点是目标节点，跳过
		if current.X == a.TargetX && current.Y == a.TargetY {
			return current, true

		}

		if a.CloseList == nil {
			a.CloseList = make(map[int]map[int]bool)
		}
		// 当前节点加入关闭列表
		if _, ok := a.CloseList[current.X]; !ok {
			a.CloseList[current.X] = make(map[int]bool)
		}
		a.CloseList[current.X][current.Y] = true

		// 遍历邻居节点
		neighbors := a.FindNeighbor(current.X, current.Y)
		for _, neighbor := range neighbors {
			// 如果邻居节点已经在关闭列表中，跳过
			// 如果是障碍物，跳过
			if a.CloseList[neighbor.X] != nil && a.CloseList[neighbor.X][neighbor.Y] {
				continue
			}
			newNode := &Node{
				X:      neighbor.X,
				Y:      neighbor.Y,
				G:      current.G + a.NeighborCost(current.X, current.Y, neighbor.X, neighbor.Y),
				H:      a.Heuristic(neighbor.X, neighbor.Y, a.TargetX, a.TargetY),
				Parent: current,
				F:      current.G + a.NeighborCost(current.X, current.Y, neighbor.X, neighbor.Y) + a.Heuristic(neighbor.X, neighbor.Y, a.TargetX, a.TargetY),
			}
			// 判断邻居节点是否在开放列表中，如果不存在，那么就加入开放列表
			// 如果存在，那么比较
			existing, index := a.OpenList.NodeInQueueu(newNode)
			if existing == nil {
				heap.Push(a.OpenList, newNode)
				continue
			}
			if existing.G > newNode.G {
				existing.G = neighbor.G
				existing.F = neighbor.F
				existing.Parent = current
				heap.Fix(a.OpenList, index)
			}
		}
	}
	return nil, false
}

// 从当前节点到目标节点的曼哈顿距离

func ManhattanDistance() Heuristic {
	return func(x1, y1, x2, y2 int) int {
		return abs(x1-x2) + abs(y1-y2)
	}
}
func FindNeighborFour() FindNeighbor {
	return func(x int, y int) []*Node {
		var ans []*Node
		var dx = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for i := 0; i < len(dx); i++ {
			X := x + dx[i][0]
			Y := y + dx[i][1]
			if X >= 0 && Y >= 0 {
				ans = append(ans, &Node{
					X: X,
					Y: Y,
				})
			}
		}
		return ans
	}
}

func ManhattanNeighborCost() NeighborCost {
	return func(x1 int, y1 int, x2 int, y2 int) int {
		return abs(x1-x2) + abs(y1-y2)
	}
}
