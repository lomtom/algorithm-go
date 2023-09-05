package _31

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

type MyCalendarTwo struct {
	time1 [][2]int
	time2 [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		time1: make([][2]int, 0),
		time2: make([][2]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	for index := range this.time1 {
		if this.time1[index][0] < end && start < this.time1[index][1] {
			return false
		}
	}
	for index := range this.time2 {
		if this.time2[index][0] < end && start < this.time2[index][1] {
			this.time1 = append(this.time1, [2]int{max(this.time2[index][0], start), min(this.time2[index][1], end)})
		}
	}
	this.time2 = append(this.time2, [2]int{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
func TestMyCalendarTwo(t *testing.T) {
	collections := []struct {
		input  [2]int
		output bool
	}{
		{
			input:  [2]int{10, 20},
			output: true,
		},
		{
			input:  [2]int{50, 60},
			output: true,
		},
		{
			input:  [2]int{10, 40},
			output: true,
		},
		{
			input:  [2]int{5, 15},
			output: false,
		},
		{
			input:  [2]int{5, 10},
			output: true,
		},
		{
			input:  [2]int{25, 55},
			output: true,
		},
	}

	s := assert.NewAssert(t)
	obj := Constructor()
	for index := range collections {
		s.Equal(collections[index].output, obj.Book(collections[index].input[0], collections[index].input[1]))
	}
}
