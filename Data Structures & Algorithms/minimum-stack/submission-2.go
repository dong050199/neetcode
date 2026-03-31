type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	res := math.MaxInt64
	for _, s := range this.stack {
		if s < res {
			res = s
		}
	}
	return res
}