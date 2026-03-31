func lastStoneWeight(stones []int) int {
	hStone := &IntHeap{}
	heap.Init(hStone)

	for _, stone := range stones {
		heap.Push(hStone, stone)
	}

	for hStone.Len() >= 2 {
		s1 := heap.Pop(hStone).(int)
		s2 := heap.Pop(hStone).(int)

		if s := smash(s1, s2);  s > 0 {
			heap.Push(hStone, s)
		}
	}

	if hStone.Len() == 1 {
		return (*hStone)[0]
	}

	return 0
}

// return -1 if both s1 and s2 destroyed after smash
func smash(s1, s2 int) int {
	if s1 == s2 {
		return -1
	}

	if s1 > s2 {
		return s1 - s2
	}

	return s2 - s1
}

// IntHeap - same as yours
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }  // Max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}