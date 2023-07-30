package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

func asteroidCollision(asteroids []int) []int {
	queue := make([]int, 0)
	compareAndPush := func(num int) {
		for num < 0 && len(queue) > 0 && queue[len(queue)-1] > 0 {
			if -num > queue[len(queue)-1] {
				queue = queue[:len(queue)-1]
			} else if -num == queue[len(queue)-1] {
				queue = queue[:len(queue)-1]
				return
			} else {
				return
			}
		}
		queue = append(queue, num)
	}
	for index := range asteroids {
		compareAndPush(asteroids[index])
	}
	return queue
}

//注意 -2, -1, 1, 2，不会发生碰撞
func TestAsteroidCollision(t *testing.T) {
	collections := []struct {
		input  []int
		output []int
	}{
		{
			[]int{5, 10, -5},
			[]int{5, 10},
		},
		{
			[]int{8, -8},
			[]int{},
		},
		{
			[]int{10, 2, -5},
			[]int{10},
		},
		{
			[]int{-2, -3, -5},
			[]int{-2, -3, -5},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{-2, -1, 1, 2},
			[]int{-2, -1, 1, 2},
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, asteroidCollision(collections[index].input))
	}
}
