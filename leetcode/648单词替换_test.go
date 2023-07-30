package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"sort"
	"strings"
	"testing"
)

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	var res []string
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	for index := range words {
		var temp string
		for dIndex := range dictionary {
			if strings.HasPrefix(words[index], dictionary[dIndex]) {
				temp = dictionary[dIndex]
				break
			}
		}
		if temp == "" {
			res = append(res, words[index])
		} else {
			res = append(res, temp)
		}
	}
	return strings.Join(res, " ")
}

func TestReplaceWords(t *testing.T) {
	type input struct {
		dictionary []string
		sentence   string
	}
	collections := []struct {
		input
		output string
	}{
		{
			input{
				[]string{"cat", "bat", "rat"},
				"the cattle was rattled by the battery",
			},
			"the cat was rat by the bat",
		},
		{
			input{
				[]string{"a", "b", "c"},
				"aadsfasf absbs bbab cadsfafs",
			},
			"a a b c",
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, replaceWords(collections[index].input.dictionary, collections[index].input.sentence))
	}
}
