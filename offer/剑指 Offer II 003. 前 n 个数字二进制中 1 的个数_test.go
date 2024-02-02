package offer

import (
	"fmt"
	"testing"
)

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

func countBits1(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func TestCountBits(t *testing.T) {
	fmt.Println(countBits(6))
	fmt.Println(countBits1(6))
}
