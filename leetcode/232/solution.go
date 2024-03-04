package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了27.86% 的Go用户
type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	stack := []int{x}
	for index := range this.stack {
		stack = append(stack, this.stack[index])
	}
	this.stack = stack
}

func (this *MyQueue) Pop() int {
	num := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return num
}

func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
