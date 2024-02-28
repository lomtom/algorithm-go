package leetcode

// [0,1,0]
// [0,2,1,0]
// [0,10,5,2]

// 执行耗时:69 ms,击败了87.65% 的Go用户
// 内存消耗:7.6 MB,击败了95.88% 的Go用户
func peakIndexInMountainArray(arr []int) int {
	var left, right = 0, len(arr)
	for left < right {
		var mid = left + (right-left)>>1
		if mid-1 >= 0 && mid+1 < len(arr) && arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
		if mid-1 >= 0 && arr[mid-1] > arr[mid] {
			right = mid - 1
		}
		if mid+1 < len(arr) && arr[mid] < arr[mid+1] {
			left = mid + 1
		}
	}
	return left
}
