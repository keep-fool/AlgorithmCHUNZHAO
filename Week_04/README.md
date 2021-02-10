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

## 实战题目

* [91. 解码方法](./decode_ways_test.go)

    [leetcode](https://leetcode-cn.com/problems/decode-ways/)
