package leetcode

// 树状数组

type NumArray struct {
	nums []int
}

func lowBit(index int) int {
	return index & -index
}

func sum(ints []int, index int) (ans int) {
	for index > 0 {
		ans += ints[index]
		index -= lowBit(index)
	}
	return ans
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	ints := make([]int, l+1)
	for index := 1; index <= l; index++ {
		index1 := index
		for index1 <= l {
			ints[index1] += nums[index-1]
			index1 += lowBit(index1)
		}
	}
	return NumArray{
		ints,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return sum(this.nums, right+1) - sum(this.nums, left)
}

// 前缀和

type NumArray1 struct {
	sums []int
}

func Constructor1(nums []int) NumArray1 {
	var sums = []int{
		nums[0],
	}
	for i := 1; i < len(nums); i++ {
		sums = append(sums, sums[i-1]+nums[i])
	}
	return NumArray1{sums: sums}
}

func (this *NumArray1) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}
