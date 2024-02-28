package leetcode

import (
	"fmt"
	"testing"
)

// 15788 11845 15024 9593 10385 10877 12360 8829  5064  5929  7660  6321  4830  7055  3761
// 764,  1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761
func Test(t *testing.T) {
	t.Log(minIncrements(15, []int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761})) // 15735
	t.Log(minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))                                                                     // 6
	t.Log(minIncrements(3, []int{5, 3, 3}))                                                                                 // 0
}

func TestBuild(t *testing.T) {
	printTree(construct([]int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761}), 0, "")
	printTree(construct([]int{1, 5, 2, 2, 3, 3, 1}), 0, "")
}

func Test1(t *testing.T) {
	for k := 0; k <= 4; k++ {
		result := (1 << k) - 1
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
