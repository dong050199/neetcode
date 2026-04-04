type LRUCache struct {
    capacity int
	cache map[int]*Node
	left *Node
	right *Node
}

type Node {
	key int
	value int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	left, right := &Node{}, &Node{}
	left.next = right
	right.prev = left
	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		left: left,
		right: right,
	}
}

func (this *LRUCache) insert(node *Node) {

}

func (this *LRUCache) remove(node *Node) {

}

func (this *LRUCache) Get(key int) int {
	if val, exist := this.cache[key]; exist {
		// remove node from existing
		// add it the right side of linked list

		return val.value
	}
}

func (this *LRUCache) Put(key int, value int) {
    
}
