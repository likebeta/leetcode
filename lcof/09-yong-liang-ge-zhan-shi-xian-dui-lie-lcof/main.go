package main

// 用两个栈实现队列
type CQueue struct {
	arr1 []int
	arr2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.arr1 = append(this.arr1, value)
}

func (this *CQueue) DeleteHead() int {
	ans := -1
	l2 := len(this.arr2)
	if l2 == 0 {
		if l2 = len(this.arr1); l2 > 0 {
			for i := l2 - 1; i >= 0; i-- {
				this.arr2 = append(this.arr2, this.arr1[i])
			}
			this.arr1 = nil
		}
	}
	if l2 > 0 {
		ans, this.arr2 = this.arr2[l2-1], this.arr2[:l2-1]
	}
	return ans
}

func main() {

}
