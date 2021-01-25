# 学习笔记

## 目录

[作业](#实战题目)

## 树

## 二叉树

## 二叉搜索树

## 堆

## 二叉堆

## 图

## 泛型递归

## 树的递归

## 实战题目

* [94.二叉树的中序遍历](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

```go
// 递归
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    inorder(root, &res)
    return res
}

func inorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    fmt.Print(root.Val)
    inorder(root.Left, res)
    *res = append(*res, root.Val)
    inorder(root.Right, res)
}
```

* [589.N叉树的前序遍历](./n_ary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)

```go
func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    results := []int{}
    for _, v := range root.Children {
        results = append(results, postorder(v)...)
    }
    results = append(results, root.Val)
    return results
}
```

* [590. N叉树的后序遍历](./n_ary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)
