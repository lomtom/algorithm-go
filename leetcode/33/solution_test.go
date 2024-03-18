package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(search([]int{5, 1, 3}, 1))             //1
	t.Log(search([]int{5, 1, 3}, 5))             //0
	t.Log(search([]int{5, 1, 3}, 3))             //2
	t.Log(search([]int{1, 3}, 3))                //1
	t.Log(search([]int{1, 3}, 2))                //-1
	t.Log(search([]int{1, 3}, 1))                //0
	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) //4
	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) //-1
	t.Log(search([]int{1}, 1))                   //0
}
