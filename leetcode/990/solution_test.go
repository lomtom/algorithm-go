package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(equationsPossible([]string{"a==b", "b!=a"}))         // false
	t.Log(equationsPossible([]string{"b==a", "a==b"}))         // true
	t.Log(equationsPossible([]string{"a==b", "b==c", "a==c"})) // true
	t.Log(equationsPossible([]string{"a==b", "b!=c", "c==a"})) // false
	t.Log(equationsPossible([]string{"c==c", "b==d", "x!=z"})) // true
	t.Log(equationsPossible([]string{"a==b", "b!=c", "c==a"})) // false
}
