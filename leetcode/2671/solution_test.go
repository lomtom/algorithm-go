package leetcode

import "testing"

func Test(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(3)
	obj.Add(2)
	obj.Add(2)
	obj.DeleteOne(3)
	obj.DeleteOne(3)
	t.Log(obj.HasFrequency(2))
}
