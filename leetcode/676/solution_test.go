package leetcode

import (
	"github.com/lomtom/go-utils/assert"
	"testing"
)

func Test(t *testing.T) {
	type input struct {
		dictionary  []string
		searchWords []string
	}
	collections := []struct {
		input
		output []bool
	}{
		{
			input{
				[]string{"hello", "leetcode"},
				[]string{
					"hello",
					"hhllo",
					"hell",
					"leetcoded",
				},
			},
			[]bool{
				false, true, false, false,
			},
		},
	}
	s := assert.NewAssert(t)
	obj := Constructor()
	for index := range collections {
		obj.BuildDict(collections[index].input.dictionary)
		for index2, value := range collections[index].input.searchWords {
			s.Equal(collections[index].output[index2], obj.Search(value))
		}
	}
}
