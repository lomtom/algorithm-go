package leetcode

func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, -n)
}

func pow(x float64, n int) float64 {
	res := 1.0
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
	}
	return res
}
