package main

import (
	"fmt"
	"sort"
)

/**
871. 最低加油次数
难度
困难

汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。
沿途有加油站，每个 station[i] 代表一个加油站，它位于出发位置东面 station[i][0] 英里处，并且有 station[i][1] 升汽油。
假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。它每行驶 1 英里就会用掉 1 升汽油。
当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。
为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。
注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。



示例 1：
输入：target = 1, startFuel = 1, stations = []才
输出：0
解释：我们可以在不加油的情况下到达目的地。

示例 2：
输入：target = 100, startFuel = 1, stations = [[10,100]]
输出：-1
解释：我们无法抵达目的地，甚至无法到达第一个加油站。

示例 3：
输入：target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
输出：2
解释：
我们出发时有 10 升燃料。
我们开车来到距起点 10 英里处的加油站，消耗 10 升燃料。将汽油从 0 升加到 60 升。
然后，我们从 10 英里处的加油站开到 60 英里处的加油站（消耗 50 升燃料）
并将汽油从 10 升加到 50 升。然后我们开车抵达目的地。
我们沿途在两个加油站停靠，所以返回 2 。
*/

type H struct {
	name string
}

func (h *H) String() string {
	return "this name" + h.name
}

func main() {
	fmt.Println(minRefuelStops(100, 50, [][]int{{25, 25}, {50, 50}}))
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	stations = append(stations, []int{target, 0}) //把终点站也当做一个油站
	result := 0
	curTotal := startFuel //当前总油量
	stock := make([]int, 0)
	for _, s := range stations {
		station := s
		for curTotal < station[0] { //如果当前总油量 小于到达下一个加油点需要的油量
			if len(stock) == 0 { //如果油储量为 0 那么不可到达
				return -1
			}
			maxS := stock[len(stock)-1] //取存油量最多的点加油
			stock = stock[:len(stock)-1]
			curTotal += maxS
			result++
		}
		stock = append(stock, station[1]) //能到到的加油站
		sort.Ints(stock)                  //排序
	}
	return result
}
