

type MinStack struct {
    stack    *linkedliststack.Stack
    minStack *linkedliststack.Stack
}

func Constructor() MinStack {
    return MinStack{
        stack:    linkedliststack.New(),
        minStack: linkedliststack.New(),
    }
}

func (this *MinStack) Push(val int) {
    this.stack.Push(val)
    minVal := val
    if !this.minStack.Empty() {
        if top, ok := this.minStack.Peek(); ok {
            if top.(int) < val {
                minVal = top.(int)
            }
        }
    }
    this.minStack.Push(minVal)
}

func (this *MinStack) Pop() {
    this.stack.Pop()
    this.minStack.Pop()
}

func (this *MinStack) Top() int {
    top, _ := this.stack.Peek()
    return top.(int)
}

func (this *MinStack) GetMin() int {
    min, _ := this.minStack.Peek()
    return min.(int)
}