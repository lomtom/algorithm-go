package leetcode

func shortestSubstrings(arr []string) []string {
	var subM = make(map[string]int)
	for index := range arr {
		var m = make(map[string]bool)
		for left := 0; left < len(arr[index]); left++ {
			for right := len(arr[index]); right >= left+1; right-- {
				m[arr[index][left:right]] = true
			}
		}
		for key := range m {
			subM[key]++
		}
	}
	var results []string
	for index := range arr {
		var result string = arr[index]
		var flag bool
		var left, count = 0, len(arr[index])
		for count > 0 {
			right := left + count
			for right <= len(arr[index]) {
				subS := arr[index][left:right]
				if v, ok := subM[subS]; ok && v == 1 && (len(subS) < len(result) || subS <= result) {
					result = subS
					flag = true
				}
				left++
				right++
			}
			count--
			left = 0
		}
		if !flag {
			result = ""
		}
		results = append(results, result)
	}
	return results
}
