package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func alertNames(keyName, keyTime []string) (ans []string) {
	timeMap := map[string][]int{}
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		timeMap[name] = append(timeMap[name], hour*60+minute)
	}
	for name, times := range timeMap {
		sort.Ints(times)
		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	sort.Strings(ans)
	return
}

func TestAlertNames(t *testing.T) {
	fmt.Println(alertNames([]string{"a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b"}, []string{"04:48", "23:53", "06:36", "07:45", "12:16", "00:52", "10:59", "17:16", "00:36", "01:26", "22:42"}))
	fmt.Println(alertNames([]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}, []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}))
	fmt.Println(alertNames([]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}, []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}))
	fmt.Println(alertNames([]string{"john", "john", "john"}, []string{"23:58", "23:59", "00:01"}))
	fmt.Println(alertNames([]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}, []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}))
}
