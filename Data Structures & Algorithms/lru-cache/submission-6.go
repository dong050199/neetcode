type LRUCache struct {
    cache map[int]*Node
    capacity int
    left *Node 
    right *Node
}

type Node struct {
    key int
    val int
    prev *Node
    next *Node
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

func (this *LRUCache) Get(key int) int {
    if node, exist := this.cache[key]; exist {
        this.remove(node)
        this.insert(node)
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist {
        node.val = value
        this.remove(node)
        this.insert(node)
        return
    }

    newNode := &Node{
        key: key,
        val: value,
    }

    this.cache[key] = newNode
    this.insert(newNode)

    for len(this.cache) > this.capacity {
        lru := this.left.next      
        this.remove(lru)
        delete(this.cache, lru.key)
    }

    return
}
