package astart

import (
	"testing"
)

func TestAStartSolution_Solution(t *testing.T) {

	astart := Solution{
		StartX:       0,
		StartY:       0,
		TargetX:      1,
		TargetY:      4,
		OpenList:     nil,
		CloseList:    nil,
		Heuristic:    ManhattanDistance(),
		FindNeighbor: FindNeighborFour(),
		NeighborCost: ManhattanNeighborCost(),
	}

	solution, b := astart.Solution()
	if b {
		way := make([][]int, 0)
		for solution != nil {
			way = append(way, []int{solution.X, solution.Y})
			solution = solution.Parent
		}
		t.Log(way)
		return
	}
	t.Log("找不到路径")
}

func TestAStartInterface(t *testing.T) {
	closeList := make(map[int]map[int]bool)
	solution := AStartSolution{
		StartNode:  &MapAStartNode{X: 0, Y: 0},
		TargetNode: &MapAStartNode{X: 1, Y: 4},
		InCloseList: func(nodeInterface AStartNode) bool {
			neighbor := nodeInterface.(*MapAStartNode)
			if closeList[neighbor.X] != nil && closeList[neighbor.X][neighbor.Y] {
				return true
			}
			return false
		},
		PutToCloseList: func(nodeInterface AStartNode) {
			current := nodeInterface.(*MapAStartNode)
			if _, ok := closeList[current.X]; !ok {
				closeList[current.X] = make(map[int]bool)
			}
			closeList[current.X][current.Y] = true
		},
		OnNodeComputer: func(node AStartNode) {
			current := node.(*MapAStartNode)
			t.Log(current.X, current.Y, current.F(), current.G(), current.H())
		},
	}
	nodeInterface, b := solution.Solution()
	if b {
		nodeInterface.PrintSolution()
	}
}
