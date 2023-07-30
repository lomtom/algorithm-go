package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

//	执行耗时:0 ms,击败了100.00% 的Go用户
//	内存消耗:2.6 MB,击败了15.63% 的Go用户
func countOfAtoms(formula string) string {
	type pair struct {
		elem string
		num  int
	}
	stack := make([]pair, 0)
	bytes := []byte(formula)
	for index := 0; index < len(bytes); {
		// 计算括号内
		if bytes[index] == ')' {
			index++
			var num = 0
			for index < len(bytes) && bytes[index] <= '9' && bytes[index] >= '0' {
				num = num*10 + int(bytes[index]-'0')
				index++
			}
			if num == 0 {
				num = 1
			}
			temp := make([]pair, 0)
			for len(stack) > 0 && stack[len(stack)-1].elem != "(" {
				stack[len(stack)-1].num = stack[len(stack)-1].num * num
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, temp...)
			continue
		}
		if bytes[index] == '(' {
			stack = append(stack, pair{
				"(",
				1,
			})
			index++
			continue
		}
		var elem string = string(bytes[index])
		var num int = 0
		index++
		// 解析原子
		for index < len(bytes) && bytes[index] <= 'z' && bytes[index] >= 'a' {
			elem += string(bytes[index])
			index++
		}
		// 解析数量
		for index < len(bytes) && bytes[index] <= '9' && bytes[index] >= '0' {
			num = num*10 + int(bytes[index]-'0')
			index++
		}
		if num == 0 {
			num = 1
		}
		stack = append(stack, pair{
			elem,
			num,
		})
	}
	m := make(map[string]int, 0)
	for index := range stack {
		m[stack[index].elem] = m[stack[index].elem] + stack[index].num
	}
	stack = make([]pair, 0)
	for k, v := range m {
		stack = append(stack, pair{
			k, v,
		})
	}
	sort.Slice(stack, func(i, j int) bool {
		return stack[i].elem < stack[j].elem
	})
	var ans string
	for index := range stack {
		ans += stack[index].elem
		if stack[index].num > 1 {
			ans += strconv.Itoa(stack[index].num)
		}
	}
	return ans
}

func TestCountOfAtoms(t *testing.T) {
	fmt.Println(countOfAtoms("H2O"))
	fmt.Println(countOfAtoms("Mg(OH)2"))
	fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
}
