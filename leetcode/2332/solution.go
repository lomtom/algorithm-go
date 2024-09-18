package leetcode

import "sort"

// 执行耗时:115 ms,击败了16.67% 的Go用户
// 内存消耗:8.9 MB,击败了75.00% 的Go用户
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	var buseIndex int
	var passengerIndex int
	var count int
	for ; buseIndex < len(buses); buseIndex++ {
		count = 0
		for count < capacity && passengerIndex < len(passengers) && buses[buseIndex] >= passengers[passengerIndex] {
			count++
			passengerIndex++
		}
	}
	ans := buses[len(buses)-1]
	passengerIndex--
	if count != capacity {
		ans = passengers[passengerIndex]
	}
	for passengerIndex >= 0 && ans == passengers[passengerIndex] {
		ans--
		passengerIndex--
	}
	return ans
}
