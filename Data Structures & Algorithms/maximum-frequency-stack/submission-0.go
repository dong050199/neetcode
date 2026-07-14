type FreqStack struct {
    freq    map[int]int
    group   map[int][]int
    maxFreq int
}

func Constructor() FreqStack {
    return FreqStack{
        freq:  make(map[int]int),
        group: make(map[int][]int),
    }
}

func (fs *FreqStack) Push(val int) {
    fs.freq[val]++
    f := fs.freq[val]

    if f > fs.maxFreq {
        fs.maxFreq = f
    }

    fs.group[f] = append(fs.group[f], val)
}

func (fs *FreqStack) Pop() int {
    stack := fs.group[fs.maxFreq]
    val := stack[len(stack)-1]

    fs.group[fs.maxFreq] = stack[:len(stack)-1]
    fs.freq[val]--

    if len(fs.group[fs.maxFreq]) == 0 {
        fs.maxFreq--
    }

    return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 