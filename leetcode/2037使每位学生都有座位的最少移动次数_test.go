package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"sort"
	"testing"
)

func minMovesToSeat(seats []int, students []int) (res int) {
	subAbs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < len(students); i++ {
		res += subAbs(seats[i], students[i])
	}
	return
}

func TestMinMovesToSeat(t *testing.T) {
	type input struct {
		seats    []int
		students []int
	}
	collections := []struct {
		input
		output int
	}{
		{
			input{
				[]int{3, 1, 5},
				[]int{2, 7, 4},
			},
			4,
		},
		{
			input{
				[]int{4, 1, 5, 9},
				[]int{1, 3, 2, 6},
			},
			7,
		},
		{
			input{
				[]int{2, 2, 6, 6},
				[]int{1, 3, 2, 6},
			},
			4,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, minMovesToSeat(collections[index].input.seats, collections[index].input.students))
	}
}
