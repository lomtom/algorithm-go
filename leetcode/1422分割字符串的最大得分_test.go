package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func maxScore(s string) int {
	one := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		}
	}
	var max int
	if s[0] == '0' {
		max = 1 + one
	} else {
		max = one - 1
	}
	temp := max
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			temp++
		} else {
			temp--
		}

		if temp > max {
			max = temp
		}
	}
	return max
}

func TestMaxScore(t *testing.T) {
	collections := []struct {
		input  string
		output int
	}{
		{
			"011101",
			5,
		},
		{
			"00111",
			5,
		},
		{
			"1111",
			3,
		},
		{
			"00",
			1,
		},
		{
			"11",
			1,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, maxScore(collections[index].input))
	}
}
