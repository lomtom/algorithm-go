package _81

import "sort"

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{map[string][]pair{}}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.m[key] = append(m.m[key], pair{timestamp, value})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	pairs := m.m[key]
	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if i > 0 {
		return pairs[i-1].value
	}
	return ""
}
