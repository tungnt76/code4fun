// https://leetcode.com/problems/implement-stack-using-queues/description/
package queue

type MyStack struct {
	arr []int
}

func Constructor() MyStack {
	return MyStack{arr: []int{}}
}

func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyStack) Pop() int {
	if len(this.arr) == 0 {
		return -1 // or handle underflow as needed
	}
	top := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return top
}

func (this *MyStack) Top() int {
	if len(this.arr) == 0 {
		return -1 // or handle underflow as needed
	}
	return this.arr[len(this.arr)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
