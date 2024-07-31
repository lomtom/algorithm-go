package leetcode

var m = map[string]byte{
	"&quot;":  '"',
	"&apos;":  '\'',
	"&gt;":    '>',
	"&lt;":    '<',
	"&frasl;": '/',
	"&amp;":   '&',
}

// 执行耗时:33 ms,击败了86.67% 的Go用户
// 内存消耗:6.9 MB,击败了40.00% 的Go用户
func entityParser(text string) string {
	var res []byte
	for i := 0; i < len(text); i++ {
		var index = i
		if text[i] == '&' {
			for j := 2; j <= 7 && i+j <= len(text); j++ {
				if v, ok := m[text[i:i+j]]; ok {
					res = append(res, v)
					i += j - 1
				}
			}
		}
		if index == i {
			res = append(res, text[i])
		}
	}
	return string(res)
}
