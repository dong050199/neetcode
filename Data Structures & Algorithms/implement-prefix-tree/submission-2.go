type PrefixTree struct {
	root *Node
}

type Node struct {
	child map[rune]*Node
	endOfWord bool
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &Node{
			child: make(map[rune]*Node),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root

	for _, w := range word {
		if _, exist := cur.child[w]; !exist {
			cur.child[w] = make(map[rune]*Node)
		}
		cur = cur.child[w]
	}

	cur.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root

}

func (this *PrefixTree) StartsWith(prefix string) bool {

}
