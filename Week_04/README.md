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

```go
func leastInterval(tasks []byte, n int) int {
    if n == 0 {
        return len(tasks)
    }
    // 将字符转切片（ASCII码 且按顺序） 并计数
    list := make([]int, 26)
    for k := range tasks {
        fmt.Println(tasks[k])
        list[tasks[k]-'A']++
    }
    fmt.Println(list)
    // maxExec 任务最大出现次数的 maxCount 拥有最大出现次数的任务的个数
    maxExec, maxCount := 0, 0
    for k := range list {
        if list[k] > maxExec {
            maxExec, maxCount = list[k], 1
        } else if list[k] == maxExec {
            maxCount++
        }
    }
    // maxExec-1 表示最大能合并连续执行的任务个数（-1 表示最后一次单独另算 无法合并组合）
    // n+1 表示间隔
    // maxCount 用来计数 maxExec-1 减1 的零头
    // 例子： 输入：tasks = ["A","A","A","B","B","B"], n = 2
    // 输出：8
    // 解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
    // maxExec-1 计 3-1 表示两轮 “A -> B -> (待命)”
    // n+1 计 2+1 表示间隔2 一轮任务就需要3
    // maxCount 计2 表示最大长度为3的 两个元素A和B
    time := (maxExec-1)*(n+1) + maxCount

    // 特殊情况 当结果集小于任务总长度时返回任务总长度，若返回结果集此时没有空余等待时间不符合题意任务无法全部调度完成
    // 输入：tasks = ["A","A","A","B","B","B", "C","C","C", "D", "D", "E"], n = 2
    if time < len(tasks) {
        return len(tasks)
    }
    return time
}
```

* [1143. 最长公共子序列](./longest_common_subsequence_test.go)

    [leetcode](https://leetcode-cn.com/problems/longest-common-subsequence/)

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([]int, len(text2)+1)
    for i := 1; i <= len(text1); i++ {
        last := 0
        for j := 1; j <= len(text2); j++ {
            tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
            if text1[i-1] == text2[j-1] {
                dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
            } else {
                dp[j] = func(a, b int) int {
                    if a > b {
                        return a
                    }
                    return b
                }(tmp, dp[j-1])
            }
            last = tmp
        }
    }
    return dp[len(text2)]
}
```

* [120. 三角形最小路径和](./triangle_test.go)

    [leetcode](https://leetcode-cn.com/problems/triangle/description/)

```go
func minimumTotal(triangle [][]int) int {
    bottom := triangle[len(triangle)-1]
    dp := make([]int, len(bottom))
    for i := range dp {
        dp[i] = bottom[i]
    }
    for i := len(dp) - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[j] = func(a, b int) int {
                if a > b {
                    return b
                }
                return a
            }(dp[j], dp[j+1]) + triangle[i][j]
        }
    }
    return dp[0]
}
```
