package leetcode

import "testing"

func Test(t *testing.T) {
	obj := Constructor(5)
	t.Log(obj.Reserve())
	t.Log(obj.Reserve())
	obj.Unreserve(2)
	t.Log(obj.Reserve())
	t.Log(obj.Reserve())
	t.Log(obj.Reserve())
	obj.Unreserve(2)
}

func Test1(t *testing.T) {
	obj := Constructor(2)
	t.Log(obj.Reserve())
	obj.Unreserve(1)
	t.Log(obj.Reserve())
	t.Log(obj.Reserve())
	obj.Unreserve(2)
	t.Log(obj.Reserve())
	obj.Unreserve(1)
	t.Log(obj.Reserve())
	obj.Unreserve(2)
	t.Log(obj.Reserve())
}
