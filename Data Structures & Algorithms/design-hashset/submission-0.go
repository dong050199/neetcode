import (
	"slices"
)

type MyHashSet struct {
	buckets  [][]int
	capacity int
	size     int
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets:  make([][]int, 8),
		capacity: 8,
		size:     0,
	}
}

func (this *MyHashSet) idx(key int) int {
	return key % this.capacity
}

func (this *MyHashSet) Add(key int) {
	curIdx := this.idx(key)
	for _, val := range this.buckets[curIdx] {
		// do nothings if idx already exist
		if key == val {
			return
		}
	}

	this.buckets[curIdx] = append(this.buckets[curIdx], key)
	this.size++
	this.resizeIfNeeded()
}

func (this *MyHashSet) resizeIfNeeded() {
	if this.size <= this.capacity*3/4  {
		return
	}

	// resize by add more storage
	this.capacity *= 2                         // x2 memory
	newBuckets := make([][]int, this.capacity) // init new memory
	for _, bucket := range this.buckets {
		for _, key := range bucket {
			newIdx := this.idx(key)
			newBuckets[newIdx] = append(newBuckets[newIdx], key)
		}
	}

	this.buckets = newBuckets
}

func (this *MyHashSet) Remove(key int) {
	idx := this.idx(key)
	for i, val := range this.buckets[idx] {
		if val == key {
			// remove val from bucket[idx]
			this.buckets[idx] = slices.Delete(this.buckets[idx], i, i+1)
            this.size--
            return
		}
	}

}

func (this *MyHashSet) Contains(key int) bool {
    idx := this.idx(key)
    for _, val := range this.buckets[idx] {
        if val == key {
            return true
        }
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */