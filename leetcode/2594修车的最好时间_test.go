package leetcode

import "math"

// 给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * n2 分钟内修好 n 辆车。
//
// 同时给你一个整数 cars ，表示总共需要修理的汽车数目。
//
// 请你返回修理所有汽车 最少 需要多少时间。
//
// 注意：所有机械工可以同时修理汽车。
// 输入：ranks = [4,2,3,1], cars = 10 输出：16
// 输入：ranks = [5,1,8], cars = 6 输出：16
func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars

	for l < r {
		m := (l + r) >> 1 // 计算中间值
		cnt := 0

		// 计算在时间 m 内可以修理的汽车数量之和
		for _, x := range ranks {
			cnt += int(math.Sqrt(float64(m / x)))
		}

		if cnt >= cars {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l) // 返回最少时间
}
