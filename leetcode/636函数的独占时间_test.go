package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	type pair struct{ idx, timestamp int }
	st := []pair{}
	for _, log := range logs {
		sp := strings.Split(log, ":")
		idx, _ := strconv.Atoi(sp[0])
		timestamp, _ := strconv.Atoi(sp[2])
		if sp[1][0] == 's' {
			if len(st) > 0 {
				ans[st[len(st)-1].idx] += timestamp - st[len(st)-1].timestamp
				st[len(st)-1].timestamp = timestamp
			}
			st = append(st, pair{idx, timestamp})
		} else {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			ans[p.idx] += timestamp - p.timestamp + 1
			if len(st) > 0 {
				st[len(st)-1].timestamp = timestamp + 1
			}
		}
	}
	return ans
}

func TestExclusiveTime(t *testing.T) {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
}
