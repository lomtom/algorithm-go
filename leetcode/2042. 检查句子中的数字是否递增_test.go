package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func areNumbersAscending(s string) bool {
	splits := strings.Split(s, " ")
	var num int
	for index := range splits {
		if '0' < splits[index][0] && splits[index][0] <= '9' {
			temp, _ := strconv.Atoi(splits[index])
			if num >= temp {
				return false
			} else {
				num = temp
			}
		}
	}
	return true
}

func TestAreNumbersAscending(t *testing.T) {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
	fmt.Println(areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"))
	fmt.Println(areNumbersAscending("4 5 11 26"))
}
