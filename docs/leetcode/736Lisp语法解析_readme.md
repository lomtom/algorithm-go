
**题目难度：**[困难](https://leetcode.cn/problems/parse-lisp-expression/)

**题目描述：**

> 给你一个类似 Lisp 语句的字符串表达式 expression，求出其计算结果。
>
> 表达式语法如下所示:
>
> - 表达式可以为整数，let 表达式，add 表达式，mult 表达式，或赋值的变量。表达式的结果总是一个整数。
> - (整数可以是正整数、负整数、0)
> - let 表达式采用 "(let v1 e1 v2 e2 ... vn en expr)" 的形式，其中 let 总是以字符串 "let"来表示，接下来会跟随一对或多对交替的变量和表达式，也就是说，第一个变量 v1被分配为表达式 e1 的值，第二个变量 v2 被分配为表达式 e2 的值，依次类推；最终 let 表达式的值为 expr表达式的值。
> - add 表达式表示为 "(add e1 e2)" ，其中 add 总是以字符串 "add" 来表示，该表达式总是包含两个表达式 e1、e2 ，最终结果是 e1 表达式的值与 e2 表达式的值之 和 。
> - mult 表达式表示为 "(mult e1 e2)" ，其中 mult 总是以字符串 "mult" 表示，该表达式总是包含两个表达式 e1、e2，最终结果是 e1 表达式的值与 e2 表达式的值之 积 。
> - 在该题目中，变量名以小写字符开始，之后跟随 0 个或多个小写字符或数字。为了方便，"add" ，"let" ，"mult" 会被定义为 "关键字" ，不会用作变量名。
> - 最后，要说一下作用域的概念。计算变量名所对应的表达式时，在计算上下文中，首先检查最内层作用域（按括号计），然后按顺序依次检查外部作用域。测试用例中每一个表达式都是合法的。有关作用域的更多详细信息，请参阅示例。

**测试用例：**

> 示例 1：
>
> 输入：expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
> 
> 输出：14
> 
> 解释：
> 
> 计算表达式 (add x y), 在检查变量 x 值时，
>
> 在变量的上下文中由最内层作用域依次向外检查。
>
> 首先找到 x = 3, 所以此处的 x 值是 3 。


> 示例 2：
> 
> 输入：expression = "(let x 3 x 2 x)"
>
> 输出：2
>
> 解释：let 语句中的赋值运算按顺序处理即可。


> 示例 3：
>
> 输入：expression = "(let x 1 y 2 x (add x y) (add x y))"
>
> 输出：5
>
> 解释：
> 
> 第一个 (add x y) 计算结果是 3，并且将此值赋给了 x 。
>
> 第二个 (add x y) 计算结果是 3 + 2 = 5 。


**限制及提示：**
> 1 <= expression.length <= 2000
>
> exprssion 中不含前导和尾随空格
>
> expressoin 中的不同部分（token）之间用单个空格进行分隔
>
> 答案和所有中间计算结果都符合 32-bit 整数范围
>
> 测试用例中的表达式均为合法的且最终结果为整数


---
**解题分析及思路：**

本题的难点在于思路以及一些细节的处理。

首先，对于一个表达式他一定是`(` + `expression` + `)`的形式，而潜逃的表达式也同样是多个此类结构的组合，那么可以将括号单独判断，直接忽略括号，重点处理里面的`expression`即可。

对于一个纯粹的`expression`，即不包括嵌套类型的表达式，他一定由变量、整数、let、add、mult组成，而let、add、mult一定是紧跟`(`，所以只要判断是`(`之后，如果立即以l、a、m开头，即可进行let、add、mult的运算。

对于let、add、mult三种运算：

1. let：可能会有以下情况：

    - 如果是变量与表达式的交替，我们需要将其变量的值以栈的形式保存起来。
    - 如果下一个字符不是小写字母，或者解析变量后下一个字符是右括号 `)`，说明是最后一个表达式，计算它的值并返回结果。并且我们需要在 scope 中清除 `let` 表达式对应的作用域变量。

2. add：后面会紧跟两个表达式，计算两个表达式的和即可
3. mult：后面紧跟两个表达式，计算两个表达式的乘积即可

对于变量和整数的解析：

- 变量：从其实位置遇到空格或者`)`即为变量的全部。
- 整数：如果为数字，临时值*10 + 该位置数字在存入临时值，遇到非数字停止解析，可能为负数，在解析钱额外判断是否以`-`开头即可。

对于括号内可能仍然以多个括号的形式，不正是和递归一致吗，所以遇到`(`进行递归即可。

**代码分析：**

解析变量：
```go
func() string {
    i0 := i
    for i < n && expression[i] != ' ' && expression[i] != ')' {
        i++
    }
    return expression[i0:i]
}
```

解析整数：
```go
func() int {
    sign, x := 1, 0
    // 判断正负
    if expression[i] == '-' {
        sign = -1
        i++
    }
    // 计算数值
    for i < n && unicode.IsDigit(rune(expression[i])) {
        x = x*10 + int(expression[i]-'0')
        i++
    }
    return sign * x
}
```

计算`let`：
```go
// 移除 "let "
i += 4
var vars []string
for {
    if !unicode.IsLower(rune(expression[i])) {
        ret = innerEvaluate() // let 表达式的最后一个 expr 表达式的值
        break
    }
    vr := parseVar()
    if expression[i] == ')' {
        vals := scope[vr]
        ret = vals[len(vals)-1] // let 表达式的最后一个 expr 表达式的值
        break
    }
    vars = append(vars, vr)
    i++ // 移除空格
    scope[vr] = append(scope[vr], innerEvaluate())
    i++ // 移除空格
}
for _, vr := range vars {
    scope[vr] = scope[vr][:len(scope[vr])-1] // 清除当前作用域的变量
}
```

计算add：
```go
// 移除 "add "
i += 4
e1 := innerEvaluate()
// 移除空格
i++
e2 := innerEvaluate()
ret = e1 + e2
```

计算mult：
```go
// 移除 "mult "
i += 5
e1 := innerEvaluate()
// 移除空格
i++
e2 := innerEvaluate()
ret = e1 * e2
```

最后代码：[Lisp语法解析](https://github.com/lomtom/algorithm-go/leetcode/736Lisp语法解析_test.go)

**复杂度：**
> 时间复杂度：O(n)
>
> 空间复杂度：O(n)

**执行结果：**
> 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
>
> 内存消耗： 2.5 MB , 在所有 Go 提交中击败了 76.19% 的用户