package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minNonZeroProduct(4))
	t.Log(minNonZeroProduct(3))
	t.Log(minNonZeroProduct(2))
}

func TestPow(t *testing.T) {
	t.Log(1e9 + 7)
	t.Log(1 << 3)
	t.Log(pow(2, 1))
	t.Log(pow(6, 2))
	t.Log(pow(14, 3))
	t.Log(pow(2, 3))
}
