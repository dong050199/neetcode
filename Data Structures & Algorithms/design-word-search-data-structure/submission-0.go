type WordDictionary struct {
    words []string
}

func Constructor() WordDictionary {
	return WordDictionary{
		words: []string{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
	this.words = append(this.words, word)
}

func (this *WordDictionary) Search(word string) bool {
	for _, w := range this.words {
		if len(w) != len(word) {
			continue
		}
		match := true
		for i := 0; i < len(word); i ++ {
			if word[i] != '.' && w[i] != word[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
