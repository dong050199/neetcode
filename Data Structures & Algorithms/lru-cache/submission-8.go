type LRUCache struct {
    cache map[int]*Node
    capacity int
    left, right *Node
}

type Node struct {
    key, value int
    prev, next *Node
}

func Constructor(capacity int) LRUCache {
    cache := make(map[int]*Node)
    left, right := &Node{}, &Node{} // create dummy node
    left.next = right
    right.prev = left
    return LRUCache{
        cache:    cache,
        capacity: capacity,
        left:     left,
        right:    right,
    }
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
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist {
        this.remove(node)
        this.insert(node)
        node.value = value
        return
    }

    newNode := &Node{
        key: key,
        value: value,
    }

    this.insert(newNode)
    this.cache[key] = newNode
    
    for len(this.cache) > this.capacity {
        cur := this.left.next
        delete(this.cache, cur.key)
        this.remove(cur)
    }
}
