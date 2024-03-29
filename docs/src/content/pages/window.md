---
title: "滑动窗口"
slug: "window"
---
说起滑动窗口，第一个能想到的是`tcp`协议中的滑动窗口（Sliding window）。

在`tcp`协议中，通常将滑动窗口用于一种流量控制。通常发送方或者接收方将数据包缓存起来，然后通过滑动窗口来控制需要发送或者接受数据包的数量或大小。
以此来改变吞吐量。

在算法中，滑动窗口也是类似的原理，需要维护一个数组/队列（可以看做tcp中缓存），然后不断的滑动窗口来计算结果。

例如[1876、长度为三且各字符不同的子字符串](../leetcode/1876长度为三且各字符不同的子字符串)

对于这题，只需要维护一个大小为3的窗口，比较窗口内的字符各不相同即可。

首先准备大小为3的窗口，即i ,i + 1,i + 2为窗口，并且对其值进行比较，如果不相等，符合条件的结果加一。
```go
func countGoodSubstrings(s string) (ans int) {
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2] {
			ans++
		}
	}
	return ans
}
```
需注意终止条件需要到`len(s)-3`，否则将会越界。

这样我们即可通过`n`的复杂度解决问题了。


简单题：

- [1876、长度为三且各字符不同的子字符串](../leetcode/substrings-of-size-three-with-distinct-characters)
- [剑指 Offer II 041、滑动窗口的平均值](../offer/qIsx9U)



参考：
- [聊聊滑动窗口](https://mp.weixin.qq.com/s/qj-XMcBw8iq5kiXItboQLA)
