package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

const mod = 1e9 + 7

func getNumberOfBacklogOrders(orders [][]int) int {
	sellsStack := make(maxPair1801, 0)
	buysStack := make(minPair1801, 0)
	for _, order := range orders {
		// buy
		if order[2] == 0 {
			for sellsStack.Len() > 0 && sellsStack[0].price <= order[0] && order[1] > 0 {
				if order[1] < sellsStack[0].amount {
					sellsStack[0].amount -= order[1]
					order[1] = 0
					break
				}
				order[1] -= heap.Pop(&sellsStack).(pair1801).amount
			}
			if order[1] > 0 {
				heap.Push(&buysStack, pair1801{order[0], order[1]})
			}
		} else {
			// sell
			for buysStack.Len() > 0 && buysStack[0].price >= order[0] && order[1] > 0 {
				if order[1] < buysStack[0].amount {
					buysStack[0].amount -= order[1]
					order[1] = 0
					break
				}
				order[1] -= heap.Pop(&buysStack).(pair1801).amount
			}
			if order[1] > 0 {
				heap.Push(&sellsStack, pair1801{order[0], order[1]})
			}
		}
	}
	sum := 0
	for index := range sellsStack {
		sum = (sum + sellsStack[index].amount) % mod
	}
	for index := range buysStack {
		sum = (sum + buysStack[index].amount) % mod
	}
	return sum
}

type pair1801 struct {
	price, amount int
}

type minPair1801 []pair1801

func (m minPair1801) Len() int {
	return len(m)
}

func (m minPair1801) Less(i, j int) bool {
	return m[i].price > m[j].price
}

func (m minPair1801) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minPair1801) Push(v interface{}) {
	*m = append(*m, v.(pair1801))
}
func (m *minPair1801) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

type maxPair1801 []pair1801

func (m maxPair1801) Len() int {
	return len(m)
}

func (m maxPair1801) Less(i, j int) bool {
	return m[i].price < m[j].price
}

func (m maxPair1801) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxPair1801) Push(v interface{}) {
	*m = append(*m, v.(pair1801))
}
func (m *maxPair1801) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

func TestGetNumberOfBacklogOrders(t *testing.T) {
	fmt.Println(
		getNumberOfBacklogOrders([][]int{
			//{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0},
			//{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1},
			//{27, 30, 0}, {10, 10, 1}, {28, 17, 1}, {19, 28, 0}, {16, 8, 1}, {14, 22, 0}, {12, 18, 1}, {3, 15, 0}, {25, 6, 1},
			{26, 7, 0}, {16, 1, 1}, {14, 20, 0}, {23, 15, 1}, {24, 26, 0}, {19, 4, 1}, {1, 1, 0},
		}))
}
