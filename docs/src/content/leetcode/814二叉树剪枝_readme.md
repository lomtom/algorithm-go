---
title: 二叉树剪枝
categories:
  - 中等
tags:
  - 二叉树
  - 分治法
number: 814
---

**题目难度：**[二叉树剪枝](https://leetcode.cn/problems/binary-tree-pruning/) 中等

**题目描述：**

> 给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
> 
> 返回移除了所有不包含 1 的子树的原二叉树。
> 
> 节点 node 的子树为 node 本身加上所有 node 的后代。

**测试用例：**

> 示例 1:
>
> ![](../img/leetcode/814二叉树剪枝/1028_2.png)
> 
> 输入：root = [1,null,0,0,1]
> 
> 输出：[1,null,0,null,1]
> 
> 解释：
> 
> 只有红色节点满足条件“所有不包含 1 的子树”。 右图为返回的答案。


> 示例 2：
> 
> ![](../img/leetcode/814二叉树剪枝/1028_1.png) 
> 
> 输入：root = [1,0,1,0,0,0,1]
> 
> 输出：[1,null,1,null,1]


> 示例 3：
>
> ![](../img/leetcode/814二叉树剪枝/1028.png)
> 
> 输入：root = [1,1,0,1,1,0,1,0]
> 
> 输出：[1,1,0,1,1,null,1]


**限制及提示：**
> 
> - 树中节点的数目在范围 [1, 200] 内
> - Node.val 为 0 或 1

---
**解题分析及思路：**

毫无疑问，这一题就是用[分治法](/dac)进行解决的了。


还是按照模板。

- 分：对于当前节点，可以将它的子节点分为两个子树。
    ```go
    // 分
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    ```
- 治：如果当前节点是nil，则返回
    ```go
    // 治
    if root == nil {
        return nil
    }
    ```
- 合：当前节点的是否被裁掉取决于当前节点的两个子节点及本身节点的值，如果当前节点的值为0，同时，左右节点均为nil，说明当前节点可以被剪掉。
    ```go
    // 合
    if root.Left== nil && root.Right == nil && root.Val == 0 {
        return nil
    }
    return root
    ```


[源代码](https://github.com/lomtom/algorithm-go/blob/main/leetcode/814二叉树剪枝_test.go)

**复杂度：**
- 时间复杂度：O(nlogn)，排序所需时间
- 空间复杂度：O(logn)，排序所需空间

**执行结果：**

- 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户 
- 内存消耗： 2.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
