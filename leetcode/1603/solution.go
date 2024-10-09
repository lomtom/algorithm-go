package leetcode

// 执行耗时:27 ms,击败了100.00% 的Go用户
// 内存消耗:7.1 MB,击败了96.88% 的Go用户

type ParkingSystem struct {
	P [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{P: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	this.P[carType-1]--
	return this.P[carType-1] >= 0
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
