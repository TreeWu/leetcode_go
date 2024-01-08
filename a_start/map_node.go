package astart

import "fmt"

var _ AStartNode = &MapAStartNode{}

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
		way = way[:len(way)-1]
	}
}

func (m *MapAStartNode) FindNeighbors() []AStartNode {
	var nodes []AStartNode
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

func (m *MapAStartNode) SetParent(nodeInterface AStartNode) {
	m.parent = nodeInterface.(*MapAStartNode)
}

func (m *MapAStartNode) Parent() AStartNode {
	return m.parent
}

func (m *MapAStartNode) Equal(nodeInterface AStartNode) bool {
	n := nodeInterface.(*MapAStartNode)
	return m.X == n.X && m.Y == n.Y
}

func (m *MapAStartNode) NeighborCost(nodeInterface AStartNode) float64 {
	n := nodeInterface.(*MapAStartNode)
	return float64(abs(m.X-n.X) + abs(m.Y-n.Y))
}

func (m *MapAStartNode) Heuristic(nodeInterface AStartNode) float64 {
	n := nodeInterface.(*MapAStartNode)
	return float64(abs(m.X-n.X) + abs(m.Y-n.Y))
}
