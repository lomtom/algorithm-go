package algorithm

import "sort"

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for index := range words {
		m[words[index]]++
	}
	tempWords := make([]string, 0, len(m))
	for s := range m {
		tempWords = append(tempWords, s)
	}

	sort.Slice(tempWords, func(i, j int) bool {
		return m[tempWords[i]] > m[tempWords[j]] || m[tempWords[i]] == m[tempWords[j]] && tempWords[i] < tempWords[j]
	})
	return tempWords[:k]
}
