# 学习笔记

## 目录

[作业](#实战题目)

[题目汇总](https://shimo.im/sheets/q9gPKWky9dj9CjdP/32ljC/)

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

* [144. 二叉树的前序遍历](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

```go
func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    preorder(root, &res)
    return res
}

func preorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    *res = append(*res, root.Val)
    preorder(root.Left, res)
    preorder(root.Right, res)
}
```

* [144. 二叉树的后序遍历](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)

```go
// 后序遍历
func postorderTraversalBinary(root *TreeNode) []int {
    res := []int{}
    postorderBinary(root, &res)
    return res
}

func postorderBinary(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    postorderBinary(root.Left, res)
    postorderBinary(root.Right, res)
    *res = append(*res, root.Val)
}
```

* [144. 二叉树的层序遍历](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

```go
func levelOrderBinary(root *TreeNode) [][]int {
    ret := [][]int{}
    if root == nil {
        return ret
    }
    q := []*TreeNode{root}
    for i := 0; len(q) > 0; i++ {
        ret = append(ret, []int{})
        p := []*TreeNode{}
        for j := 0; j < len(q); j++ {
            node := q[j]
            ret[i] = append(ret[i], node.Val)
            if node.Left != nil {
                p = append(p, node.Left)
            }
            if node.Right != nil {
                p = append(p, node.Right)
            }
        }
        q = p
    }
    return ret
}
```

* [236. 二叉树的最近公共祖先](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root == q || root == p {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
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

```go
// 中序遍历 递归
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

* [最小的k个数](./n_ary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

```go
// 直接排序 返回前K个
func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    return arr[:k]
}
```
