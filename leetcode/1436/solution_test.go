package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(destCity([][]string{
		{"Lima", "Sao Paulo"},
		{"London", "New York"},
		{"New York", "Lima"},
	}))
	t.Log(destCity([][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}))
	t.Log(destCity([][]string{
		{"A", "Z"},
	}))
}
