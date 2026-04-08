type LRUCache struct {
    cache map[int]*Node
    cap int
    left, right *Node
}

type Node struct {
    key int
    value int
    prev, next *Node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        cap: capacity,
        cache: make(map[int]*Node),
        left: &Node{},
        right: &Node{},
    }

    lru.left.next = lru.right
    lru.right.prev = lru.left 
    return lru
}

func (this *LRUCache) remove(node *Node) {
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
}

func (this *LRUCache) insert(node *Node) {
    prev, next := this.right.prev, this.right
    node.prev = prev
    node.next = next
    prev.next = node
    next.prev = node
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
        delete(this.cache, key)
    } 

    newNode := &Node{
        key: key, 
        value: value,
    }

    this.cache[key] = newNode
    this.insert(newNode)

    if len(this.cache) > this.cap {
        lru := this.left.next
        this.remove(lru)
        delete(this.cache, lru.key)
    }
}
