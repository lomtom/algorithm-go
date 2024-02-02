---
title: 四叉树交集
categories:
  - 中等
tags:
  - 分治法
  - 树
number: 558
---
**题目难度：**[中等](https://leetcode.cn/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/)

**题目描述：**

> 二进制矩阵中的所有元素不是 0 就是 1 。
>
> 给你两个四叉树，quadTree1 和 quadTree2。其中 quadTree1 表示一个 n * n 二进制矩阵，而 quadTree2 表示另一个 n * n 二进制矩阵。
>
> 请你返回一个表示 n * n 二进制矩阵的四叉树，它是 quadTree1 和 quadTree2 所表示的两个二进制矩阵进行 按位逻辑或运算 的结果。
>
> 注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
>
> 四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
>
> - val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
> - isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。

```go
type Node struct {
Val         bool
IsLeaf      bool
TopLeft     *Node
TopRight    *Node
BottomLeft  *Node
BottomRight *Node
}
```

> 我们可以按以下步骤为二维区域构建四叉树：
>
> - 如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
> - 如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
> - 使用适当的子网格递归每个子节点。

![](../img/leetcode/558四叉树交集/new_top.png)


> 四叉树格式：
>
> 输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节点。
>
> 它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。
>
> 如果 isLeaf 或者 val 的值为 True ，则表示它在列表 [isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为 False ，则表示值为 0 。


**测试用例：**

> 示例 1:
>
> ![](../img/leetcode/558四叉树交集/qt1.png)
>
> ![](../img/leetcode/558四叉树交集/qt2.png)
>
> 输入：quadTree1 = [[0,1],[1,1],[1,1],[1,0],[1,0]]
>
> , quadTree2 = [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
>
> 输出：[[0,0],[1,1],[1,1],[1,1],[1,0]]
>
> 解释：quadTree1 和 quadTree2 如上所示。由四叉树所表示的二进制矩阵也已经给出。
>
> 如果我们对这两个矩阵进行按位逻辑或运算，则可以得到下面的二进制矩阵，由一个作为结果的四叉树表示。
>
> 注意，我们展示的二进制矩阵仅仅是为了更好地说明题意，你无需构造二进制矩阵来获得结果四叉树。
>
> ![](../img/leetcode/558四叉树交集/qtr.png)


> 示例 2:
>
> 输入：quadTree1 = [[1,0]]
>
> , quadTree2 = [[1,0]]
>
> 输出：[[1,0]]
>
> 解释：两个数所表示的矩阵大小都为 1*1，值全为 0
>
> 结果矩阵大小为 1*1，值全为 0 。


> 示例 3：
> 输入：quadTree1 = [[0,0],[1,0],[1,0],[1,1],[1,1]]
>
> , quadTree2 = [[0,0],[1,1],[1,1],[1,0],[1,1]]
>
> 输出：[[1,1]]

> 示例 4：
>
> 输入：quadTree1 = [[0,0],[1,1],[1,0],[1,1],[1,1]]
>
> , quadTree2 = [[0,0],[1,1],[0,1],[1,1],[1,1],null,null,null,null,[1,1],[1,0],[1,0],[1,1]]
>
> 输出：[[0,0],[1,1],[0,1],[1,1],[1,1],null,null,null,null,[1,1],[1,0],[1,0],[1,1]]


> 示例 5：
>
> 输入：quadTree1 = [[0,1],[1,0],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
>
> , quadTree2 = [[0,1],[0,1],[1,0],[1,1],[1,0],[1,0],[1,0],[1,1],[1,1]]
>
> 输出：[[0,0],[0,1],[0,1],[1,1],[1,0],[1,0],[1,0],[1,1],[1,1],[1,0],[1,0],[1,1],[1,1]]


**限制及提示：**
> - quadTree1 和 quadTree2 都是符合题目要求的四叉树，每个都代表一个 n * n 的矩阵。
> - n == 2^x ，其中 0 <= x <= 9.

---
**解题分析及思路：**

乍一看，题目都没看懂。

简单来说，就是一棵四个节点的树放在一个小格子里，如果当前节点拥有四个子节点，那么四个子节点将再次瓜分这个格子为四个小格子。

然后求，两棵树各自形成的小格子做逻辑或运算，最终将结果保存到同样的四叉树中并返回。

这个逻辑或运算是当前两棵树相同位置的值的或运算。

题目讲解完毕，那就是怎么来计算了。

对于这样的树的计算，很适合使用[分治法](/dac)。

- 分：将两棵树当前节点的四个子节点拆分来。即`quadTree1.TopLeft`对应`quadTree2.TopLeft`，`quadTree1.TopRight`对应 `quadTree2.TopRight`等
- 治：判断当前节点是否满足临界终止，并进行计算返回。即当前的两棵树所访问到的当前节点有一个是叶子节点时，终止分的操作，通过计算 `quadTree1`|`quadTree2`操作来进行计算并返回值
- 合：通过两棵树的四个子节点计算结果来计算当前节点的值，即是当前节点的最后结果。即两棵树的四个子节点的计算结果会影响当前节点的值，达到一定条件改变当前节点值。

**代码分析：**

将所有计算的结果都放在`quadTree1`这棵树上，最终返回这棵树便是最终结果，这样无需重新初始化新的树。

- 分的操作：将两棵树的四个节点进行拆分，并将最终结果分别放到`quadTree1`的四个子节点上。
```go
quadTree1.TopLeft = intersect(quadTree1.TopLeft, quadTree2.TopLeft)
quadTree1.TopRight = intersect(quadTree1.TopRight, quadTree2.TopRight)
quadTree1.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
quadTree1.BottomRight = intersect(quadTree1.BottomRight, quadTree2.BottomRight)
```

- 治的操作：判断当前访问到的节点是否为叶子节点，如果有一方为叶子节点，则根据 `quadTree1.Val`| `quadTree2.Val`的值来判断最终返回的节点。
```go
if quadTree1.IsLeaf {
    if quadTree1.Val {
        return quadTree1
    }
    return quadTree2
}
if quadTree2.IsLeaf {
    return intersect(quadTree2, quadTree1)
}
```

- 合的操作：先假设当前节点为非叶子节点，设置`quadTree1.IsLeaf = false`,如果四个子节点都是叶子节点并且值都相同，则将当前节点设为叶子节点，并改变当前节点值以及置空四个叶子节点。
```go
quadTree1.IsLeaf = false
quadTree1.Val = false
if quadTree1.TopLeft.IsLeaf && quadTree1.TopRight.IsLeaf && quadTree1.BottomLeft.IsLeaf && quadTree1.BottomRight.IsLeaf && quadTree1.TopLeft.Val == quadTree1.TopRight.Val && quadTree1.TopRight.Val == quadTree1.BottomLeft.Val && quadTree1.BottomLeft.Val == quadTree1.BottomRight.Val {
    quadTree1.Val = quadTree1.TopLeft.Val
    quadTree1.IsLeaf = true
    quadTree1.TopLeft = nil
    quadTree1.TopRight = nil
    quadTree1.BottomLeft = nil
    quadTree1.BottomRight = nil
}
return quadTree1
```

源代码：[四叉树交集](https://github.com/lomtom/algorithm-go/blob/main/leetcode/558四叉树交集_test.go)

**复杂度：**

- 时间复杂度：O(n ^ 2)
- 空间复杂度：O(1)

**执行结果：**

- 执行用时： 8 ms , 在所有 Go 提交中击败了 92.86% 的用户
- 内存消耗： 6.5 MB , 在所有 Go 提交中击败了 50.00% 的用户
