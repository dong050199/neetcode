type KthLargest struct {
	k int 
	h *MinHeap
}


func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}

	for _, num := range nums {
		heap.Push(h, num)
	}

	for h.Len() > k {
		heap.Pop(h)
	}
 

	return KthLargest{
		k: k,
		h: h,
	}
}


func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
	} else if val > (*this.h)[0]  {
		heap.Pop(this.h)
		heap.Push(this.h, val)
	}
    return (*this.h)[0]
}

type MinHeap []int
 
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
 
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
 
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
 
