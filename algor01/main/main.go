package main

import (
	gk "GoMyself/algor01/GreedyKnapsack"
	"fmt"
)

func main() {
	gk.Capacity = 10
	var w []int = []int{1, 2, 3, 4, 5}
	var v []int = []int{3, 10, 6, 3, 5}
	var ks *gk.Ks = gk.InitKS(w, v)

	ks.ComputeRatio()
	ks.SortRatio()
	profit, xi := ks.ComputeProfit()
	fmt.Println("背包的最大收益:")
	fmt.Println(profit)
	fmt.Println("最大收益情况下各种物品放入背包的对应比例:([物品]比例)")
	fmt.Println(xi)
}
