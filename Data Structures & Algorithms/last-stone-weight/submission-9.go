func lastStoneWeight(stones []int) int {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for maxHeap.Len() >= 2 {
		first := heap.Pop(maxHeap).(int)
		second := heap.Pop(maxHeap).(int)

		switch {
			case first == second:
				continue
			case first > second:
				heap.Push(maxHeap, first - second)
			default:
				heap.Push(maxHeap, second - first)
		}
	}

	if maxHeap.Len() == 0 {
		return 0
	}

	return (*maxHeap)[0]
}


type IntHeap []int
    
    func (h IntHeap) Len() int           { return len(h) }
    func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
    func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
    
    func (h *IntHeap) Push(x interface{}) {
        // Push and Pop use pointer receivers because they modify the slice's length,
        // not just its contents.
        *h = append(*h, x.(int))
    }
    
    func (h *IntHeap) Pop() interface{} {
        old := *h
        n := len(old)
        x := old[n-1]
        *h = old[0 : n-1]
        return x
    }
    
    // This example inserts several ints into an IntHeap, checks the minimum,
    // and removes them in order of priority.
    // func Example_intHeap() {
    //     h := &IntHeap{2, 1, 5}
    //     heap.Init(h)
    //     heap.Push(h, 3)
    //     fmt.Printf("minimum: %d\n", (*h)[0])
    //     for h.Len() > 0 {
    //         fmt.Printf("%d ", heap.Pop(h))
    //     }
    //     // Output:
    //     // minimum: 1
    //     // 1 2 3 5
    // }