# 学习笔记

## 目录

[作业](#实战题目)

[题目汇总](https://shimo.im/sheets/q9gPKWky9dj9CjdP/32ljC/)

## 实战题目

* [169. 多数元素](./majority_element_test.go)

    [leetcode](https://leetcode-cn.com/problems/majority-element/)

```go
// 暴力
func majorityElement(nums []int) int {
    maps := make(map[int]int)
    for _, v := range nums {
        maps[v]++
    }
    var out, max int
    for k, v := range maps {
        if v > max {
            max = v
            out = k
        }
    }
    return out
}

// 投票法
func majorityElement1(nums []int) int {
    major := 0
    count := 0

    for _, num := range nums {
        if count == 0 {
            major = num
        }
        if major == num {
            count += 1
        } else {
            count -= 1
        }
    }

    return major
}
```
