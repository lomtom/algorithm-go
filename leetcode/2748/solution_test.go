package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(countBeautifulPairs([]int{756, 1324, 2419, 495, 106, 111, 1649, 1474, 2001, 1633, 273, 1804, 2102, 1782, 705, 1529, 1761, 1613, 111, 186, 412}))
	t.Log(countBeautifulPairs([]int{1799, 259, 1453, 374, 1854, 2212, 2104, 2221}))
	t.Log(countBeautifulPairs([]int{51, 65, 87, 6, 17, 32, 14, 42, 46}))
	t.Log(countBeautifulPairs([]int{68, 8, 84, 14, 88, 42, 53}))
	t.Log(countBeautifulPairs([]int{2, 5, 1, 4}))
	t.Log(countBeautifulPairs([]int{11, 21, 12}))
}

func Test1(t *testing.T) {
	var count int
	for _, value := range m {
		count += len(value)
	}
	t.Log(count)

	var count1 int
	for i := 1; i <= 9; i++ {
		t.Logf("\nkey：%d \t", i)
		for j := 1; j <= 9; j++ {
			if gcd(i, j) {
				count1++
				t.Logf("%d \t", j)
			}
		}
	}
	t.Log(count1)
}

func gcd(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}
	return a == 1
}
