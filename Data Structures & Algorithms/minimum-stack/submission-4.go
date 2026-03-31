type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// update min stack
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
		return
	}

	if val > this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
	// update min stack
	this.minStack = this.minStack[:len(this.minStack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
