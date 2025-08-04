package leetcode

// 执行耗时:293 ms,击败了42.24% 的Go用户
// 内存消耗:94.3 MB,击败了13.37% 的Go用户

type WordDictionary struct {
	WordDictionaries [26]*WordDictionary
	IsEnd            bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		WordDictionaries: [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.IsEnd = true
		return
	}
	now := word[0] - 'a'
	if this.WordDictionaries[now] == nil {
		n := Constructor()
		this.WordDictionaries[now] = &n
	}
	this.WordDictionaries[now].AddWord(word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.IsEnd
	}
	if word[0] == '.' {
		for index := range this.WordDictionaries {
			if v := this.WordDictionaries[index]; v != nil {
				if v.Search(word[1:]) {
					return true
				}
			}
		}
	} else {
		now := word[0] - 'a'
		if v := this.WordDictionaries[now]; v != nil {
			return v.Search(word[1:])
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
