package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

func TestName(t *testing.T) {
	// true waterbottle erbottlewat
	// false
	fmt.Println(isFlipedString("abcd", "acdb"))
}
