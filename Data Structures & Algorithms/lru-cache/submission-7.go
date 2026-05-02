type LRUCache struct {
    cache map[int]*Node
    capacity int
    left *Node
    right *Node
}

type Node struct {
    key int
    value int
    next *Node
    prev *Node
}

func (this *LRUCache) insert(node *Node) {
    prev, next := this.right.prev, this.right
    node.next = next
    node.prev = prev
    prev.next = node
    next.prev = node
}

func (this *LRUCache) remove(node *Node) {
    prev, next := node.prev, node.next
    prev.next = next 
    next.prev = prev
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        cache: make(map[int]*Node),
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
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist {
        this.remove(node)
        this.insert(node)
        this.cache[key].value = value
        return
    }

    node := &Node{
        key: key,
        value: value,
    }

    this.insert(node)
    this.cache[key] = node

    for len(this.cache) > this.capacity {
        cur := this.left.next
        this.remove(this.left.next)
        delete(this.cache, cur.key)
    }
    return
}
