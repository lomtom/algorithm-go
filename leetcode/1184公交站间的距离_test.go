package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	l := len(distance)
	ans := []int{0, 0}
	for i := 0; i < l; i++ {
		if i >= start && i < destination {
			ans[0] += distance[i]
		} else {
			ans[1] += distance[i]
		}
	}
	if ans[0] > ans[1] {
		return ans[1]
	}
	return ans[0]
}
