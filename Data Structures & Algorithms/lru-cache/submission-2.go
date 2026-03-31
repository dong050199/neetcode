type LRUCache struct {
    capacity int
	cache map[int]*Node
	left *Node
	right *Node
}

type Node struct {
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
	prev, next := this.right.prev , this.right
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if val, exist := this.cache[key]; exist {
		this.remove(val)
		this.insert(val)
		return val.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, exist := this.cache[key]; exist {
		this.remove(val)
		this.insert(val)
		val.value = value
		return
	}

	node := &Node{
		key: key,
		value: value,
	}

	this.insert(node)
	this.cache[key] = node

	if len(this.cache) > this.capacity {
		// evic the least recenlty use from left
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
