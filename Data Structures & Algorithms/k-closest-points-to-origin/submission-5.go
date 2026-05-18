func kClosest(points [][]int, k int) [][]int {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for _, point := range points {
		p := Point{
			point: point,
			distance: point[0] * point[0]  + point[1] * point[1],
		}

		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}

	res := [][]int{}
	for maxHeap.Len() > 0 {
		res = append(res, heap.Pop(maxHeap).(Point).point)
	}

	return res
}

type Point struct {
	point []int
	distance int
}

type IntHeap []Point

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].distance > h[j].distance } // min heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(Point))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}