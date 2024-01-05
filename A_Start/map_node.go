package main

import "fmt"

var _ AStartNodeInterface = &MapAStartNode{}

type MapAStartNode struct {
	X, Y int
	DefaultNode
	parent *MapAStartNode
}

func (m *MapAStartNode) PrintSolution() {
	var way []*MapAStartNode
	for m != nil {
		way = append(way, m)
		m = m.parent
	}
	for len(way) != 0 {
		fmt.Printf("(%d,%d)->", way[len(way)-1].X, way[len(way)-1].Y)
	}
}

func (m *MapAStartNode) FindNeighbors() []AStartNodeInterface {
	var nodes []AStartNodeInterface
	var dx = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < len(dx); i++ {
		X := m.X + dx[i][0]
		Y := m.Y + dx[i][1]
		if X >= 0 && Y >= 0 {
			nodes = append(nodes, &MapAStartNode{
				X: X,
				Y: Y,
			})
		}
	}
	return nodes
}

func (m *MapAStartNode) SetParent(nodeInterface AStartNodeInterface) {
	m.parent = nodeInterface.(*MapAStartNode)
}

func (m *MapAStartNode) Parent() AStartNodeInterface {
	return m.parent
}

func (m *MapAStartNode) Equal(nodeInterface AStartNodeInterface) bool {
	n := nodeInterface.(*MapAStartNode)
	return m.X == n.X && m.Y == n.Y
}

func (m *MapAStartNode) NeighborCost(nodeInterface AStartNodeInterface) float64 {
	n := nodeInterface.(*MapAStartNode)
	return float64(abs(m.X-n.X) + abs(m.Y-n.Y))
}

func (m *MapAStartNode) Heuristic(nodeInterface AStartNodeInterface) float64 {
	n := nodeInterface.(*MapAStartNode)
	return float64(abs(m.X-n.X) + abs(m.Y-n.Y))
}
