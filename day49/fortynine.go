package day49

/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
*/

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &WordDictionary{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *WordDictionary) searchInNode(word string, node *WordDictionary) bool {
	for i, ch := range word {
		if ch == '.' {
			for _, child := range node.children {
				if child != nil && this.searchInNode(word[i+1:], child) {
					return true
				}
			}
			return false
		} else {
			ch -= 'a'
			if node.children[ch] == nil {
				return false
			}
			node = node.children[ch]
		}
	}
	return node != nil && node.isEnd
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchInNode(word, this)
}
