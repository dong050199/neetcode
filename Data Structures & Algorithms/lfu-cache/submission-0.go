type LFUCache struct {
	cache    map[int]*Node
	list     map[int]*List
	capacity int
	min      int
}

type List struct {
	right *Node
	left  *Node
	size  int
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
	freq int
}

func (l *List) insert(node *Node) {
	prev, next := l.right.prev, l.right
	node.prev = prev
	node.next = next
	prev.next = node
	next.prev = node
	l.size++
}

func (l *List) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
	l.size--
}

func makeList() *List {
	left := &Node{}
	right := &Node{}
	left.next = right
	right.prev = left
	return &List{
		left:  left,
		right: right,
	}
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*Node),
		list:     make(map[int]*List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Get(key int) int {
    node, exist := this.cache[key]
    if !exist {
        return -1
    }

    // remove node from frequency list
    list, exist := this.list[node.freq]
    if exist {
        list.remove(node)
    }

    node.freq++
    nextList, exist := this.list[node.freq]
    if !exist {
        nextList = makeList()
    }

    nextList.insert(node)
    this.list[node.freq] = nextList

    if list.size == 0 && this.min == node.freq-1 {
        this.min++
    }

    return node.val
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 {
        return 
    }

    node, exist := this.cache[key]
    if exist {
        node.val = value
        this.Get(key)
        return
    }

    if len(this.cache) >= this.capacity {
        minList := this.list[this.min]
        lfuNode := minList.left.next
        minList.remove(lfuNode)
        delete(this.cache, lfuNode.key)
    }

    newNode := &Node{
        key: key,
        val: value,
        freq: 1,
    }

    this.min = 1
    list, exist := this.list[newNode.freq]
    if !exist {
        list = makeList()
    }

    list.insert(newNode)
    this.list[newNode.freq] = list
    this.cache[key] = newNode
}