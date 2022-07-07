package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

type MyCalendar struct {
	time [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		time: make([][2]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for index := range this.time {
		if this.time[index][0] <= start && start < this.time[index][1] || this.time[index][0] < end && end < this.time[index][1] || start <= this.time[index][0] && this.time[index][1] <= end {
			return false
		}
	}
	this.time = append(this.time, [2]int{start, end})
	return true
}

//   10 - 20
//1. 20 - 30 true
//2. 11 - 15 false
//3. 10 - 15 false
//4. 15 - 20 false
//5. 8  - 20 false
//6. 8  - 22 false
//7. 8  - 10 true
//8. 8  - 15 false
//9. 15 - 22 false

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

//[47,50],[33,41],[39,45],[33,42],[25,32],[26,35],[19,25],[3,8],[8,13],[18,27]
//[true,  true,	  false,  false,  true,    false,  true,  true,  true,  false]
func TestMyCalendar(t *testing.T) {
	collections := []struct {
		input  [2]int
		output bool
	}{
		{
			input:  [2]int{47, 50},
			output: true,
		},
		{
			input:  [2]int{33, 41},
			output: true,
		},
		{
			input:  [2]int{39, 45},
			output: false,
		},
		{
			input:  [2]int{33, 42},
			output: false,
		},
		{
			input:  [2]int{25, 32},
			output: true,
		},
		{
			input:  [2]int{26, 35},
			output: false,
		},
		{
			input:  [2]int{19, 25},
			output: true,
		},
		{
			input:  [2]int{3, 8},
			output: true,
		},
		{
			input:  [2]int{8, 13},
			output: true,
		},
		{
			input:  [2]int{18, 27},
			output: false,
		},
	}

	s := assert.NewAssert(t)
	obj := Constructor()
	for index := range collections {
		s.Equal(collections[index].output, obj.Book(collections[index].input[0], collections[index].input[1]))
	}
}
