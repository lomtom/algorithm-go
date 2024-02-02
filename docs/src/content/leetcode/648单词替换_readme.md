---
title: 单词替换
categories:
  - 中等
tags:
  - 数组
number: 648
---
**题目难度：**[中等](https://leetcode.cn/problems/replace-words/)

**题目描述：**

> 在英语中，我们有一个叫做 词根(root) 的概念，可以词根后面添加其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
> 
> 现在，给定一个由许多词根组成的词典 dictionary 和一个用空格分隔单词形成的句子 sentence。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
>
> 你需要输出替换之后的句子。


**测试用例：**

> 示例1 ：
>
> 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
> 
> 输出："the cat was rat by the bat"


> 示例 2：
>
> 输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
> 
> 输出："a a b c"

**限制及提示：**
> 1 <= dictionary.length <= 1000
> 
> 1 <= dictionary[i].length <= 100
> 
> dictionary[i] 仅由小写字母组成。
> 
> 1 <= sentence.length <= 10^6
> 
> sentence 仅由小写字母和空格组成。
> 
> sentence 中单词的总量在范围 [1, 1000] 内。
> 
> sentence 中每个单词的长度在范围 [1, 1000] 内。
> 
> sentence 中单词之间由一个空格隔开。
> 
> sentence 没有前导或尾随空格。


---
**解题分析及思路：**

本题本身没有什么难度，数据量又少，几乎可以不考虑时间以及空间上的事情。

唯一需要注意的是怎么在一个单词中（例如`another`）找出它的最短词根（例如`an`），可以考虑使用`strings`工具包中的`strings.HasPrefix`方法进行判断。

找到最短词根：那么，我们直接遍历`sentence`分割后的字符数组，判断`dictionary`是否含有词根，再将其与临时变量比较长度，如果长度小于临时变量，即替换临时变量，这样就可以找到最短词根了。

优化：为了减少判断次数，可以事先将`dictionary`进行排序，可以使用`sort.Slice`方法对`dictionary`根据长度进行排序。这样，在判断是否含有词根时，如果找到了，直接停止比较。


**代码分析：**

分割sentence
```go
words := strings.Split(sentence, " ")
```

将`dictionary`按照长度进行排序，短的排前面
```go
sort.Slice(dictionary, func(i, j int) bool {
    return len(dictionary[i]) < len(dictionary[j])
})
```

遍历、比较以及保存最后结果
```go
var res []string
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
```



最后代码：[单词替换](https://github.com/lomtom/algorithm-go/blob/main/leetcode/648单词替换_test.go)

**复杂度：**

> 时间复杂度：O(d * s)，d为`dictionary`的长度，s为`sentence`的字符数
>
> 空间复杂度：O(s)，s为`sentence`的字符数

**执行结果：**

> 执行用时： 8 ms , 在所有 Go 提交中击败了 100.00% 的用户
>
> 内存消耗： 8.4 MB , 在所有 Go 提交中击败了 84.68% 的用户
