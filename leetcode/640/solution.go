package leetcode

import (
	"fmt"
	"strings"
)

func solveEquation(equation string) string {
	splits := strings.Split(equation, "=")

	cal := func(s string) (coefficient, number int) {
		temps := strings.Replace(s, "-", "+-", -1)
		ss := strings.Split(temps, "+")
		for index := range ss {
			if ss[index] == "" {
				continue
			}
			temp, sign, num, hasx, valid := ss[index], 1, 0, false, false
			for i := range temp {
				if temp[i] == '-' {
					sign = -1
					continue
				}
				if temp[i] == 'x' {
					hasx = true
					continue
				}
				valid = true
				num = num*10 + int(temp[i]-'0')
			}
			sum := sign
			if valid {
				sum *= num
			}
			if hasx {
				coefficient += sum
			} else {
				number += sum
			}
		}
		return
	}
	coefficient1, number1 := cal(splits[0])
	coefficient2, number2 := cal(splits[1])
	if coefficient1 == coefficient2 {
		if number1 == number2 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return fmt.Sprintf("x=%v", (number2-number1)/(coefficient1-coefficient2))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：1.8 MB, 在所有 Go 提交中击败了78.95%的用户
