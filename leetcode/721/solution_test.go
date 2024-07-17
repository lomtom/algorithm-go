package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(accountsMerge([][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"}}))
	t.Log(accountsMerge([][]string{
		{"David", "David0@m.co", "David1@m.co"},
		{"David", "David3@m.co", "David4@m.co"},
		{"David", "David4@m.co", "David5@m.co"},
		{"David", "David2@m.co", "David3@m.co"},
		{"David", "David1@m.co", "David2@m.co"}}))
}

//测试用例:[
//["David","David0@m.co","David1@m.co"],
//["David","David3@m.co","David4@m.co"],
//["David","David4@m.co","David5@m.co"],
//["David","David2@m.co","David3@m.co"],
//["David","David1@m.co","David2@m.co"]]
//测试结果:[["David","David0@m.co"],["David","David1@m.co","David2@m.co","David3@m.co","David4@m.co","David5@m.co"]]
//期望结果:[["David","David0@m.co","David1@m.co","David2@m.co","David3@m.co","David4@m.co","David5@m.co"]]
