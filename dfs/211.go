package dfs

// 211.添加与搜索单词-数据结构设计
type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.root.Insert(word)
}

func (t *TrieNode) Insert(s string) {
	node := t
	for _, ch := range s {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(this.root, word)
}

func (this *WordDictionary) dfs(node *TrieNode, word string) bool {
	if node == nil {
		return false
	}
	if word == "" {
		return node.isEnd
	}
	ch := word[0]
	if ch != '.' {
		return this.dfs(node.children[ch-'a'], word[1:])
	} else {
		for _, child := range node.children {
			if this.dfs(child, word[1:]) {
				return true
			}
		}
	}
	return false
}
