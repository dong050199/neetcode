func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}

	if len(stones) == 1 {
		return stones[0]
	}


	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() >= 2 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		newStone := a - b
		newStone = abs(newStone)
		heap.Push(h, newStone)
	}

	return (*h)[0]
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return 0 - i
}


type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap: i > j
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}