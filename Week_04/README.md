# 学习笔记

## 目录

[作业](#实战题目)

[题目汇总](https://shimo.im/sheets/q9gPKWky9dj9CjdP/32ljC/)

## 动态规划(Dynamic Programming)

> simplifying a complicated problem by breaking it down into simpler sub-problems in a recursive manner  
> 将一个复杂无问题分解为简单的子问题(用递归的方式)

与递归分治无根本区别

同：找到重复子问题
异：DP需要有最优子结构 中途淘汰次优解

> 注: 动态规划 分治都是运营递归的算法  都需要分解复杂问题为子问题。

* 分治即没有最优子结构  所有问题都要计算一遍，最终合并结果（传统意义上成为 分治 分而治之）

### 关键点

1. 最优子结构: opt[n] = best_of(opt[n-1],opt[n-2],...)
2. 存储中间状态: opt[i]
3. 递推公式(状态转移方程 DP方程)
    * Fib:opt[n-1]+opt[n-2]
    * 二维路径:opt[i,j]=opt[i+1][j]+opt[i][j+1](且判断边界 和 a[i,j]是否为空地)

## 实战题目

* [91. 解码方法](./decode_ways_test.go)

    [leetcode](https://leetcode-cn.com/problems/decode-ways/)

```go
func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    pre, cur := 1, 1
    for i := 1; i < len(s); i++ {
        switch {
        case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
            return 0
        case s[i] == '0':
            cur = pre
        case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
            pre, cur = cur, cur+pre
        default:
            pre = cur
        }
    }
    return cur
}
```

* [64. 最小路径和](./minimum_path_sum_test.go)

    [leetcode](https://leetcode-cn.com/problems/minimum-path-sum/)

```go
func minPathSum(grid [][]int) int {
    if grid == nil {
        return 0
    }
    // 初始化DP
    dp := make([][]int, len(grid))
    for k, v := range grid {
        dp[k] = make([]int, len(v))
    }
    dp[0][0] = grid[0][0]
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            switch {
            case i == 0 && j != 0:
                dp[i][j] = dp[i][j-1] + grid[i][j]
            case i != 0 && j == 0:
                dp[i][j] = dp[i-1][j] + grid[i][j]
                // 最优
            case i != 0 && j != 0:
                dp[i][j] = func(a, b int) int {
                    if a > b {
                        return b
                    }
                    return a
                }(dp[i][j-1], dp[i-1][j]) + grid[i][j]
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}
```

* [621. 任务调度器](./task_scheduler_test.go)

    [leetcode](https://leetcode-cn.com/problems/task-scheduler/)
