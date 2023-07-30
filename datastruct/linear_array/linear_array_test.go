package linear_array

import (
	"fmt"
	"testing"
)

// lowBit 取值：取最低位1及后面所有的0所表达的数（a[i]所管理的数的数量）
// 例如：
// X -> 二进制-> 取值后->
// 1 -> 01   -> 1    -> 1
// 2 -> 10   -> 10   -> 2
// 3 -> 11   -> 1    -> 1
// 4 -> 100  -> 100  -> 4
// 5 -> 101  -> 1    -> 1
// 6 -> 110  -> 10   -> 2
// 7 -> 111  -> 1    -> 1
// 8 -> 1000 -> 1000 -> 8
func lowBit(index int) int {
	// 原码 & 补码（原码按位取反加一）
	// 6 -> 110 & 010 -> 10 -> 2
	return index & -index
}

// 在a[i]加上k，只需更新所有a[i]的上级节点
func add(index, k int, a []int) {
	l := len(a)
	for index <= l {
		a[index] += k
		index += lowBit(index)
	}
}

// 求a[1]到a[i]的和，只需求a[i]及其之前的峰的和
func sum(index int, a []int) (ans int) {
	for index >= 1 {
		ans += a[index]
		index -= lowBit(index)
	}
	return
}

// 初始化当前值的同时，将自身的值累加到父亲节点
func initTree(nums []int) []int {
	l := len(nums)
	a := make([]int, l+1)
	for i := 0; i < l; i++ {
		index := i + 1
		a[index] += nums[i]
		fatherIndex := index + lowBit(index)
		if fatherIndex <= l {
			a[fatherIndex] += a[index]
		}
	}
	return a
}

func TestLowBit(t *testing.T) {
	collections := []int{1, 2, 3, 4, 5, 6, 7, 8, 16, 32, 33, 34, 35, 88, 89, 90, 91}
	for _, value := range collections {
		fmt.Printf("value: %d\tlowbit: %d\n", value, lowBit(value))
	}
}

func TestAdd(t *testing.T) {
	array := make([]int, 9)
	add(1, 1, array)
	fmt.Println(array)
	add(4, 2, array)
	fmt.Println(array)
	add(5, 3, array)
	fmt.Println(array)
}

func TestSum(t *testing.T) {
	array := []int{0, 1, 3, 3, 10, 5, 11, 7, 36}
	for index := 1; index <= 8; index++ {
		fmt.Printf("index: %d\tsum: %d\n", index, sum(index, array))
	}
}

func TestInit(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(initTree(array))
}
