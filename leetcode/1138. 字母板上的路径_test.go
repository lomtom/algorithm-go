package leetcode

import (
	"fmt"
	"testing"
)

func alphabetBoardPath(target string) string {
	now := [2]int{0, 0}
	var result []byte
	m := make(map[byte][2]int)
	for i := 0; i < 26; i++ {
		m[byte('a'+i)] = [2]int{i / 5, i % 5}
	}
	for index := range target {
		coordinate := m[target[index]]
		if coordinate[0] < now[0] {
			for i := coordinate[0]; i < now[0]; i++ {
				result = append(result, 'U')
			}
		}
		if coordinate[1] < now[1] {
			for i := coordinate[1]; i < now[1]; i++ {
				result = append(result, 'L')
			}
		}
		if coordinate[0] > now[0] {
			for i := now[0]; i < coordinate[0]; i++ {
				result = append(result, 'D')
			}
		}
		if coordinate[1] > now[1] {
			for i := now[1]; i < coordinate[1]; i++ {
				result = append(result, 'R')
			}
		}
		result = append(result, '!')
		now = coordinate
	}
	return string(result)
}

func TestAlphabetBoardPath(t *testing.T) {
	fmt.Println(alphabetBoardPath("leet"))
	fmt.Println(alphabetBoardPath("code"))
}
