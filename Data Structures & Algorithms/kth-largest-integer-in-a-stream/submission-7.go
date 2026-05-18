type KthLargest struct {
    h *IntHeap
    k int
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
        h: h,
        k: k,
    }
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)

    if this.h.Len() > this.k {
        heap.Pop(this.h)
    }

    return (*this.h)[0]
}

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