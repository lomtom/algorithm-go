package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
	"unicode"
)

func evaluate(expression string) int {
	i, n := 0, len(expression)
	// 解析变量
	parseVar := func() string {
		i0 := i
		for i < n && expression[i] != ' ' && expression[i] != ')' {
			i++
		}
		return expression[i0:i]
	}
	// 解析整数
	parseInt := func() int {
		sign, x := 1, 0
		// 判断正负
		if expression[i] == '-' {
			sign = -1
			i++
		}
		// 计算数值
		for i < n && unicode.IsDigit(rune(expression[i])) {
			x = x*10 + int(expression[i]-'0')
			i++
		}
		return sign * x
	}
	scope := map[string][]int{}
	// 定义方法
	var innerEvaluate func() int
	innerEvaluate = func() (ret int) {
		// 非表达式，可能为：整数或变量
		if expression[i] != '(' {
			// 变量
			if unicode.IsLower(rune(expression[i])) {
				vals := scope[parseVar()]
				return vals[len(vals)-1]
			}
			// 整数
			return parseInt()
		}
		// 移除左括号
		i++
		// "let" 表达式
		if expression[i] == 'l' {
			// 移除 "let "
			i += 4
			var vars []string
			for {
				if !unicode.IsLower(rune(expression[i])) {
					ret = innerEvaluate() // let 表达式的最后一个 expr 表达式的值
					break
				}
				vr := parseVar()
				if expression[i] == ')' {
					vals := scope[vr]
					ret = vals[len(vals)-1] // let 表达式的最后一个 expr 表达式的值
					break
				}
				vars = append(vars, vr)
				i++ // 移除空格
				scope[vr] = append(scope[vr], innerEvaluate())
				i++ // 移除空格
			}
			for _, vr := range vars {
				scope[vr] = scope[vr][:len(scope[vr])-1] // 清除当前作用域的变量
			}
		} else if expression[i] == 'a' { // "add" 表达式
			// 移除 "add "
			i += 4
			e1 := innerEvaluate()
			// 移除空格
			i++
			e2 := innerEvaluate()
			ret = e1 + e2
		} else { // "mult" 表达式
			// 移除 "mult "
			i += 5
			e1 := innerEvaluate()
			// 移除空格
			i++
			e2 := innerEvaluate()
			ret = e1 * e2
		}
		// 移除右括号
		i++
		return
	}
	return innerEvaluate()
}

func TestEvaluate(t *testing.T) {
	collections := []struct {
		input  string
		output int
	}{
		{
			"(let x 2 (mult x (let x 3 y 4 (add x y))))",
			14,
		},
		{
			"(let l 3 l 2 l)",
			2,
		},
		{
			"(let x 1 y 2 x (add x y) (add x y))",
			5,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(evaluate(collections[index].input), collections[index].output)
	}
}
