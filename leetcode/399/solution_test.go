package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(calcEquation([][]string{{"a", "b"}, {"a", "c"}, {"d", "e"}, {"d", "f"}, {"a", "d"}, {"aa", "bb"}, {"aa", "cc"}, {"dd", "ee"}, {"dd", "ff"}, {"aa", "dd"}, {"a", "aa"}},
		[]float64{2.0, 3.0, 4.0, 5.0, 7.0, 5.0, 8.0, 9.0, 3.0, 2.0, 2.0}, [][]string{{"ff", "a"}})) //0.08333
	t.Log(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
	t.Log(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
	t.Log(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}))
}

//
//输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
//输出：[3.75000,0.40000,5.00000,0.20000]
//示例 3：
//
//输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
//输出：[0.50000,2.00000,-1.00000,-1.00000]
