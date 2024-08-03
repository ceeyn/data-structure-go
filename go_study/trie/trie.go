package trie

type Node struct {
	isWord bool
	next   map[rune]*Node
}

func newNode(isWord bool) *Node {
	return &Node{
		isWord: isWord,
		next:   make(map[rune]*Node),
	}
}

type Trie struct {
	root  *Node
	count int
}

func NewTrie() *Trie {
	trie := Trie{
		root:  newNode(false),
		count: 0,
	}
	return &trie
}

func (t *Trie) Insert(word string) {
	curNode := t.root
	for i := 0; i < len(word); i++ {
		var c rune = rune(word[i])
		next, exits := curNode.next[c]
		if !exits {
			curNode.next[c] = newNode(false)
		}
		curNode = next
	}
	if !curNode.isWord {
		curNode.isWord = true
		t.count++
	}
}

func (t *Trie) Contains(word string) bool {
	curNode := t.root
	for _, v := range word {
		if _, ex := curNode.next[v]; !ex {
			return false
		}
		curNode = curNode.next[v]
	}
	return true
}
func (t *Trie) IsPrefix(word string) bool {
	cur := t.root
	for _, v := range word {
		if _, ex := cur.next[v]; !ex {
			return false
		}
		cur = cur.next[v]
	}
	return true
}
