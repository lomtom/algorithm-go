package _636

import (
	"github.com/lomtom/go-utils/assert"
	"sort"
	"testing"
)

// 执行耗时:4 ms,击败了87.50% 的Go用户
// 内存消耗:3.4 MB,击败了17.50% 的Go用户
func frequencySort(nums []int) (ans []int) {
	m := make(map[int]int)
	for index := range nums {
		m[nums[index]]++
	}
	type pair struct {
		num   int
		count int
	}
	var temp []pair
	for k, v := range m {
		temp = append(temp, pair{
			k, v,
		})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].count < temp[j].count || temp[i].count == temp[j].count && temp[i].num > temp[j].num
	})
	for index := range temp {
		for i := 0; i < temp[index].count; i++ {
			ans = append(ans, temp[index].num)
		}
	}
	return
}

func TestFrequencySort(t *testing.T) {

	collections := []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 1, 2, 2, 2, 3},
			[]int{3, 1, 1, 2, 2, 2},
		},
		{
			[]int{2, 3, 1, 3, 2},
			[]int{1, 3, 3, 2, 2},
		},
		{
			[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			[]int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, frequencySort(collections[index].input))
	}
}
