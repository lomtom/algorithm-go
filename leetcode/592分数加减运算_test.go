package leetcode

import (
	"fmt"
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

//-1/2+1/2+1/3
func fractionAddition(expression string) string {
	l := len(expression)
	denominator, numerator := 0, 1 // 分子，分母
	for i := 0; i < l; {
		sign := 1
		// 读取符号
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		// 读取分子
		denominator1 := 0
		for i < l && expression[i] != '/' {
			denominator1 = denominator1*10 + int(expression[i]-'0')
			i++
		}
		denominator1 *= sign
		i++
		// 读取分母
		numerator1 := 0
		for i < l && expression[i] != '+' && expression[i] != '-' {
			numerator1 += numerator1*10 + int(expression[i]-'0')
			i++
		}
		denominator = denominator*numerator1 + denominator1*numerator
		numerator *= numerator1
	}
	if denominator == 0 {
		return "0/1"
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	g := gcd(abs(denominator), numerator)
	return fmt.Sprintf("%d/%d", denominator/g, numerator/g)
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：1.9 MB, 在所有 Go 提交中击败了60.00%的用户
func TestFractionAddition(t *testing.T) {
	collections := []struct {
		input  string
		output string
	}{
		{"-5/2+10/3+7/9", "29/18"},
		{"-1/2+1/2", "0/1"},
		{"-1/2+1/2+1/3", "1/3"},
		{"1/3-1/2", "-1/6"},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, fractionAddition(collections[index].input))
	}
}
