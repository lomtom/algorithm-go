package _206

type Skiplist struct {
	nums map[int]int
}

func Constructor() Skiplist {
	return Skiplist{
		make(map[int]int),
	}
}

func (this *Skiplist) Search(target int) bool {
	if count, ok := this.nums[target]; ok && count > 0 {
		return true
	}
	return false
}

func (this *Skiplist) Add(num int) {
	this.nums[num]++
}

func (this *Skiplist) Erase(num int) bool {
	if count, ok := this.nums[num]; ok && count > 0 {
		this.nums[num]--
		return true
	}
	return false
}
