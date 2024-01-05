package main

import (
	"testing"
)

func TestAStartSolution_Solution(t *testing.T) {

	astart := Solution{
		StartX:       0,
		StartY:       0,
		TargetX:      4,
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
		for solution.Parent != nil {
			way = append(way, []int{solution.X, solution.Y})
			solution = solution.Parent
		}
		for i := 0; i <= len(way)/2; i++ {
			way[i], way[len(way)-1-i] = way[len(way)-1-i], way[i]
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
		TargetNode: &MapAStartNode{X: 4, Y: 4},
		InCloseList: func(nodeInterface AStartNodeInterface) bool {
			neighbor := nodeInterface.(*MapAStartNode)
			if closeList[neighbor.X] != nil && closeList[neighbor.X][neighbor.Y] {
				return true
			}
			return false
		},
		PutToCloseList: func(nodeInterface AStartNodeInterface) {
			current := nodeInterface.(*MapAStartNode)
			if _, ok := closeList[current.X]; !ok {
				closeList[current.X] = make(map[int]bool)
			}
			closeList[current.X][current.Y] = true
		},
	}
	nodeInterface, b := solution.Solution()
	if b {
		nodeInterface.PrintSolution()
	}
}
