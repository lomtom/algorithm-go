package leetcode

import (
	"math"
)

// 执行耗时:3 ms,击败了27.87% 的Go用户
// 内存消耗:2.2 MB,击败了95.08% 的Go用户
func maxHeightOfTriangle(red int, blue int) int {
	blueLayer1, blueLayer2 := findNumberOfTerms(1, blue), findNumberOfTerms(2, blue)
	redLayer1, redLayer2 := findNumberOfTerms(1, red), findNumberOfTerms(2, red)
	return max(min(blueLayer1*2-1, redLayer2*2), min(blueLayer2*2, redLayer1*2-1)) + 1
}

// 定义求解等差数列项数的方法
func findNumberOfTerms(a, Sn int) int {
	// 等差数列求和公式: Sn = (n / 2) * (2a + (n - 1) * b)
	// 我们需要解出 n, 根据公式重写为二次方程：
	// n^2 * b + n * (2a - b) - 2 * Sn = 0
	// n^2 * 2 + n * (2a - 2) - 2 * Sn = 0
	// n^2 + n * (a - 1) - Sn = 0

	// 计算二次方程的系数
	A := float64(1)
	B := float64(a - 1)
	C := float64(-Sn)

	// 使用求解二次方程的公式：n = (-B ± sqrt(B^2 - 4AC)) / 2A
	discriminant := B*B - 4*A*C

	// 计算两个根
	n1 := (-B + math.Sqrt(discriminant)) / (2 * A)
	// n2 := (-B - math.Sqrt(discriminant)) / (2 * A) 必为负数，舍弃

	// 我们只需要正整数解
	return int(n1)
}
