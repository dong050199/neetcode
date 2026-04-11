func kClosest(points [][]int, k int) [][]int {
	// composition of distance + point
	h := &DtHeap{}

	for _, point := range points {
		heap.Push(h, PointWithDistance{
			point: point,
			distance: distance(point),
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := [][]int{}

	for k > 0 {
		res = append(res, heap.Pop(h).(PointWithDistance).point)
		k--
	}

	return res
}

type PointWithDistance struct {
	point []int
	distance float64
}

func distance(point []int) float64 {
	x := point[0]*point[0] + point[1]*point[1]
	return math.Sqrt(float64(x))
}


type DtHeap []PointWithDistance

func (h DtHeap) Len() int           { return len(h) }
func (h DtHeap) Less(i, j int) bool { return h[i].distance < h[j].distance } // Min-heap property
func (h DtHeap) Swap(i, j int)      { h[i].distance, h[j].distance = h[j].distance, h[i].distance }

func (h *DtHeap) Push(x any) {
	*h = append(*h, x.(PointWithDistance))
}

func (h *DtHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
