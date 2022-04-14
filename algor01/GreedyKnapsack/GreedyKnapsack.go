package greedyknapsack

var Capacity int   // 背包容量
var profit float64 // 物品收益

// var xi map[int]float64 //最大收益情况下各种物品放入背包的对应比例
// 初始化对象
type knapsack struct {
	index  int
	weight int
	value  int
	profit float64
}

type Ks []*knapsack

// type greedyknapsacker interface {
// }

func InitKS(int_weight, int_value []int) *Ks {
	ks := make(Ks, 5)
	for i := 0; i < 5; i++ {
		// ks[i].weight = int_weight[i]
		// ks[i].value = int_value[i]
		// ks[i].profit = 0
		ks[i] = &knapsack{index: i + 1, weight: int_weight[i], value: int_value[i], profit: 0}
	}
	return &ks
}

// func (ks ks) initgk(int_weight, int_value []int) {

// 	for i := 0; i < len(int_weight); i++ {
// 		ks[i].weight = int_weight[i]
// 		ks[i].value = int_value[i]
// 		ks[i].profit = 0
// 	}
// }

func (ks *Ks) ComputeRatio() {
	r := *ks
	for i := 0; i < 5; i++ {

		w := (r[i]).weight
		v := (r[i]).value
		index := (r[i]).index
		var p float64 = float64(v) / float64(w)
		r[i] = &knapsack{index: index, weight: w, value: v, profit: p}
	}

}

// 根据单位重量的收益值比大小，对物品进行排序
func (ks *Ks) SortRatio() {
	r := *ks
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if (r[j]).profit > (r[i]).profit {
				temp := r[i]
				r[i] = r[j]
				r[j] = temp
			}
		}
	}
}

// 按背包物品index排序

// 计算背包问题的最优解和所放物品的最大收益
func (ks *Ks) ComputeProfit() (float64, map[int]float64) {
	var temp int = 0
	var i int = 0
	r := *ks
	xi := make(map[int]float64, 5)
	for {
		if !(temp <= Capacity) {
			break
		}
		if i == 5 {
			break
		} else {
			if (r[i].weight + temp) <= Capacity {
				profit += (float64)(r[i].value)
				temp += r[i].weight
				xi[r[i].index] = float64(r[i].weight) / float64(Capacity)
			} else {
				var _weight int = Capacity - temp
				profit += (float64)(_weight) / (float64)(r[i].profit)
				temp += _weight
				xi[r[i].index] = float64(_weight) / float64(Capacity)
			}
		}
		i++
	}
	return profit, xi
}


