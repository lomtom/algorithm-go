package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lowBit(index int) int {
	return index & -index
}

func add(index, num int, tree []int) {
	l := len(tree)
	for index < l {
		tree[index] += num
		index += lowBit(index)
	}
}

func sum(index int, tree []int) (ans int) {
	for index >= 1 {
		ans += tree[index]
		index &= index - 1
	}
	return ans
}

func Init(array []int) []int {
	l := len(array)
	ans := make([]int, l+1)
	for i := 0; i < l; i++ {
		index := i + 1
		ans[index] += array[i]
		fatherIndex := index + lowBit(index)
		if fatherIndex <= l {
			ans[fatherIndex] += ans[index]
		}
	}
	return ans
}

func main() {
	p130(os.Stdin, os.Stdout)
}

func p130(_r io.Reader, _w io.Writer) {
	input, output := bufio.NewReader(_r), bufio.NewWriter(_w)
	var n, q int
	fmt.Fscan(input, &n, &q)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &array[i])
	}
	tree := Init(array)
	for i := 0; i < q; i++ {
		operate := [3]int{}
		fmt.Fscan(input, &operate[0], &operate[1], &operate[2])
		if operate[0] == 1 {
			add(operate[1], operate[2], tree)
		} else if operate[0] == 2 {
			ans := sum(operate[2], tree) - sum(operate[1]-1, tree)
			fmt.Fprintln(output, ans)
		}
	}
	output.Flush()
}
