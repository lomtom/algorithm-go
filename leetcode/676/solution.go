package leetcode

// 执行用时：8 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：6.8 MB, 在所有 Go 提交中击败了74.42%的用户

type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		make(map[int][]string),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for index := range dictionary {
		this.m[len(dictionary[index])] = append(this.m[len(dictionary[index])], dictionary[index])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	l := len(searchWord)
	if values, ok := this.m[l]; ok {
		for _, value := range values {
			var count int
			for index := 0; index < l; index++ {
				if value[index] != searchWord[index] {
					count++
				}
			}
			if count == 1 {
				return true
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
