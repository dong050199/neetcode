type KthLargest struct {
    k    int
    nums *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    heap.Init(h)  
    
    for _, num := range nums {
        heap.Push(h, num) 
    }
    
    for h.Len() > k {
        heap.Pop(h)  
    }
    
    return KthLargest{
        k:    k,
        nums: h,
    }
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this.nums, val)
    // because heap also can append to the existing node?
    if this.nums.Len() > this.k {
        heap.Pop(this.nums)
    }
    
    return (*this.nums)[0]  // Return root (kth largest)
}

// IntHeap - same as yours
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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