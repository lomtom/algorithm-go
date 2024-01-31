package lcr

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

// [1,3,1]
// [1,5,1]
// [4,2,1]

// 执行耗时:5 ms,击败了43.14% 的Go用户
// 内存消耗:3.3 MB,击败了96.08% 的Go用户
func jewelleryValue(frame [][]int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for indexI := range frame {
		for indexJ := range frame[indexI] {
			if indexI == 0 && indexJ == 0 {
				continue
			}
			if indexI == 0 {
				frame[indexI][indexJ] = frame[indexI][indexJ] + frame[indexI][indexJ-1]
				continue
			}
			if indexJ == 0 {
				frame[indexI][indexJ] = frame[indexI][indexJ] + frame[indexI-1][indexJ]
				continue
			}
			frame[indexI][indexJ] = frame[indexI][indexJ] + max(frame[indexI-1][indexJ], frame[indexI][indexJ-1])
		}
	}
	return frame[len(frame)-1][len(frame[0])-1]
}

func TestJewelleryValue(t *testing.T) {
	collections := []struct {
		input  [][]int
		output int
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			12,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, jewelleryValue(collections[index].input))
	}
}
