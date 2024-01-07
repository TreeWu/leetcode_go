package astart

import (
	"fmt"
	"math"
)

var _ AStartNode = &PlaneTicketNode{}

type PlaneTicketNode struct {
	Airport string //机场名称
	DefaultNode
	parent    *PlaneTicketNode
	Cost      map[string]float64          //到达各个机场的花费
	Neighbors map[string]*PlaneTicketNode //到达各个机场的机票
}

func (p *PlaneTicketNode) FindNeighbors() []AStartNode {
	var as []AStartNode
	for k := range p.Neighbors {
		as = append(as, p.Neighbors[k])
	}
	return as
}

func (p *PlaneTicketNode) SetParent(node AStartNode) {
	p.parent = node.(*PlaneTicketNode)
}

func (p *PlaneTicketNode) Parent() AStartNode {

	return p.parent
}

func (p *PlaneTicketNode) Equal(node AStartNode) bool {
	return p.Airport == node.(*PlaneTicketNode).Airport
}

func (p *PlaneTicketNode) NeighborCost(node AStartNode) float64 {
	return p.Cost[node.(*PlaneTicketNode).Airport]
}

func (p *PlaneTicketNode) Heuristic(node AStartNode) float64 {
	if cost, ok := p.Cost[node.(*PlaneTicketNode).Airport]; ok {
		return cost
	}
	return math.MaxFloat64
}
func (p *PlaneTicketNode) SetCost(plane *PlaneTicketNode, cost float64) {
	if p.Cost == nil {
		p.Cost = make(map[string]float64)
	}
	if p.Neighbors == nil {
		p.Neighbors = make(map[string]*PlaneTicketNode)
	}
	p.Cost[plane.Airport] = cost
	p.Neighbors[plane.Airport] = plane
}

func (p *PlaneTicketNode) PrintSolution() {

	var sum float64
	var way []*PlaneTicketNode
	cur := p
	for p != nil {
		way = append(way, p)
		cur = p
		p = p.parent
		if p != nil {
			sum += p.Cost[cur.Airport]
		}
	}
	fmt.Println("总花费:", sum)
	for len(way) != 0 {
		fmt.Printf("%s->", way[len(way)-1].Airport)
		way = way[:len(way)-1]
	}
}
