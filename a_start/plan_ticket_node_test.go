package astart

import "testing"

func TestPlaneTicketNode(t *testing.T) {
	maps := make(map[string]*PlaneTicketNode)
	maps["北京"] = &PlaneTicketNode{Airport: "北京"}
	maps["天津"] = &PlaneTicketNode{Airport: "天津"}
	maps["河南"] = &PlaneTicketNode{Airport: "河南"}
	maps["深圳"] = &PlaneTicketNode{Airport: "深圳"}
	maps["重庆"] = &PlaneTicketNode{Airport: "重庆"}

	maps["北京"].SetCost(maps["天津"], 100)
	maps["北京"].SetCost(maps["河南"], 100)

	maps["天津"].SetCost(maps["深圳"], 40)
	maps["天津"].SetCost(maps["河南"], 100)

	maps["河南"].SetCost(maps["深圳"], 100)
	maps["河南"].SetCost(maps["重庆"], 100)

	maps["深圳"].SetCost(maps["重庆"], 40)

	closeList := make(map[string]struct{})

	solution := AStartSolution{
		StartNode:  maps["北京"],
		TargetNode: maps["重庆"],
		OpenList:   nil,
		InCloseList: func(node AStartNode) bool {
			neighbor := node.(*PlaneTicketNode)
			if _, ok := closeList[neighbor.Airport]; ok {
				return true
			}
			return false
		},
		PutToCloseList: func(node AStartNode) {
			neighubor := node.(*PlaneTicketNode)
			closeList[neighubor.Airport] = struct{}{}
		},
		OnNodeComputer: func(node AStartNode) {
			current := node.(*PlaneTicketNode)
			t.Log(current.Airport, current.F(), current.G(), current.H())
		},
	}
	node, b := solution.Solution()
	if b {
		node.PrintSolution()
	}

}
