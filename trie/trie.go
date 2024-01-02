package trie

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	t := Trie{NewTrieNode()}
	return &t
}

type TrieNode struct {
	Value      string
	Children   map[string]*TrieNode
	IsTerminal bool
}

func NewTrieNode() *TrieNode {
	t := TrieNode{}
	t.Children = make(map[string]*TrieNode)
	return &t
}

func (t *Trie) Insert(item string) {
	p := t.Root
	for i := 0; i < len(item); i++ {
		c := string(item[i])
		next, ok := p.Children[c]
		if !ok {
			next = NewTrieNode()
			p.Children[c] = next
		}
		p = next
	}
	p.Value = item
	p.IsTerminal = true
}

func (t *Trie) Search(item string) (bool, bool) {
	p := t.Root
	for i := 0; i < len(item); i++ {
		c := string(item[i])
		next, ok := p.Children[c]
		if !ok {
			return false, false
		}
		p = next
	}
	return true, p.IsTerminal
}
