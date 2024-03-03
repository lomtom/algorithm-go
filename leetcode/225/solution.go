package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了75.48% 的Go用户
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	queue := []int{x}
	for index := range this.queue {
		queue = append(queue, this.queue[index])
	}
	this.queue = queue
}

func (this *MyStack) Pop() int {
	pop := this.queue[0]
	this.queue = this.queue[1:]
	return pop
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
