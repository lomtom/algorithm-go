package leetcode

import "testing"

func Test(t *testing.T) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	t.Log(myStack.Top())   // 返回 2
	t.Log(myStack.Pop())   // 返回 2
	t.Log(myStack.Empty()) // 返回 False
}
