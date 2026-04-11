type KthLargest struct {
	k int
	stream []int
}


func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k: k,
		stream: nums,
	}
}


func (this *KthLargest) Add(val int) int {
	// brute force approach
	this.stream = append(this.stream, val)
	sort.Slice(this.stream[:], func(i, j int) bool {
    	return this.stream[i] > this.stream[j] 
	})
	return this.stream[this.k-1]
}
