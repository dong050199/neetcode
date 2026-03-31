type MinStack struct {
	stack []int
	minStack []int
	top int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
		top: -1,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val) // push to stack
	// update min stack and top
	if this.top == -1 {
		this.minStack = append(this.minStack, val)
		this.top++
		return
	}
	// if the top of minStack is lessen then val so we should keep the lessen one.
	if this.minStack[this.top] >= val {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[this.top])
	}
	this.top++
	return
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.minStack = this.minStack[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.stack[this.top]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.top]
}
