package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

//执行用时：8 ms, 在所有 Go 提交中击败了67.76%的用户
//内存消耗：6.2 MB, 在所有 Go 提交中击败了32.57%的用户
func frequencySort(s string) string {
	m := make(map[uint8]int)
	for index := range s {
		m[s[index]]++
	}
	res := make([][]uint8, len(s)+1)
	for key, value := range m {
		if res[value] == nil {
			res[value] = []uint8{key}
		} else {
			res[value] = append(res[value], key)
		}
	}
	var ans []uint8
	for i := len(res) - 1; i > 0; i-- {
		if res[i] == nil || len(res[i]) == 0 {
			continue
		}
		for _, elem := range res[i] {
			for count := 0; count < i; count++ {
				ans = append(ans, elem)
			}
		}
	}
	return string(ans)
}

func TestFrequencySort(t *testing.T) {
	collections := []struct {
		input  string
		output string
	}{
		{
			"tree",
			"eetr",
		},
		{"cccaaa",
			"cccaaa",
		},
		{
			"Aabb",
			"bbAa",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, frequencySort(collections[index].input))
	}
}
