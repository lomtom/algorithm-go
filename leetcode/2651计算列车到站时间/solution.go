package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了14.29% 的Go用户
func TestFindDelayedArrivalTime(t *testing.T) {
	type input = struct {
		arrivalTime int
		delayedTime int
	}
	collections := []struct {
		input  input
		output int
	}{
		{
			input{
				15,
				5,
			},
			20,
		},
		{
			input{
				13,
				11,
			},
			0,
		},
		{
			input{
				23,
				24,
			},
			23,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, findDelayedArrivalTime(collections[index].input.arrivalTime, collections[index].input.delayedTime))
	}
}
