package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// 执行耗时:48 ms,击败了93.94% 的Go用户
// 内存消耗:9.8 MB,击败了93.94% 的Go用户
func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}

func TestRemoveSubfolders(t *testing.T) {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
