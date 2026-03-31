type Node struct {
	children map[rune]*Node
	endOfWord bool
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &Node{
			children: make(map[rune]*Node),
			endOfWord: false,
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if _, exist := cur.children[c]; !exist {
			cur.children[c] = &Node{
				children: make(map[rune]*Node),
			}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, w := range word {
		if _, exist := cur.children[w]; !exist {
			return false
		}
		cur = cur.children[w]
	}
	return cur.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, w := range prefix {
		if _, exist := cur.children[w]; !exist {
			return false
		}
		cur = cur.children[w]
	}
	return true
}
