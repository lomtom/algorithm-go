package offer

import "strings"

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
