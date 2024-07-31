package minstack

type MinStack struct {
	data [][]int
}

func Constructor() MinStack {
	return MinStack{
		data: make([][]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	newMin := val
	if len(this.data) > 0 {
		newMin = min(this.GetMin(), val)
	}

	newTop := []int{val, newMin}
	this.data = append(this.data, newTop)
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1][1]
}

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(val);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */
