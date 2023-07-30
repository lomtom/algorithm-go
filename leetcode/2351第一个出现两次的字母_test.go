package leetcode

import (
	"fmt"
	"testing"
)

func repeatedCharacter(s string) byte {
	var res int
	for index := range s {
		temp := 1 << (s[index] - 'a')
		if res&temp > 0 {
			return s[index]
		}
		res |= temp
	}
	return ' '
}

func TestRepeatedCharacter(t *testing.T) {
	fmt.Printf("%b\n", 1<<0)
	fmt.Printf("%b\n", 1<<1)
	fmt.Printf("%b\n", 1<<2)
	fmt.Printf("%b\n", 1<<3)
	fmt.Println(repeatedCharacter("abcbaacz"))
}
