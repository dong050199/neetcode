type LRUCache struct {
    cache map[int]*Node
    capacity int
    left, right *Node
}

type Node struct {
    key int
    value int
    next *Node
    prev *Node
}

func (this *LRUCache)insert(node *Node) {
    prev, next := this.right.prev, this.right
    prev.next = node
    next.prev = node
    node.next = next
    node.prev = prev
}

func (this *LRUCache)remove(node *Node) {
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
}

func Constructor(capacity int) LRUCache {
    cache := make(map[int]*Node)
    lru := LRUCache{
        cache: cache,
        capacity: capacity,
        left: &Node{},
        right: &Node{},
    }
    lru.left.next = lru.right
    lru.right.prev = lru.left
    return lru
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.cache[key]; exist {
        this.remove(node)
        this.insert(node)
        return this.cache[key].value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        delete(this.cache, key)
    }

    node := &Node{
        key: key, 
        value: value,
    }

    this.cache[key] = node
    this.insert(node)

    if len(this.cache) > this.capacity {
        lru := this.left.next
        this.remove(lru)
        delete(this.cache, lru.key)
    }
}
