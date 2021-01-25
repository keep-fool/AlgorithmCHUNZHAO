# 学习笔记

## 目录

## 数组 数组(Array)

* 简述

* 时间复杂度
  * prepend O(1)
  * append O(1)
  * lookup O(1)
  * insert O(n)
  * delete O(n)

* 优劣

  * 优点：支持随机访问，在知道下标的情况下时间复杂度为O(1), 不知道下标时，用二分法查找，时间复杂度为O(logn)
  * 缺点:低效的插入与删除
    假设数组的长度为 n，现在，如果我们需要将一个数据插入到数组中的第 k 个位置。为了把第 k 个位置腾出来，给新来的数据，我们需要将第 k～n 这部分的元素都顺序地往后挪一位。时间复杂度为(1+2+...n)/n = O(n)
    如果数组中存储的数据并没有任何规律，数组只是被当作一个存储数据的集合。在这种情况下，如果要将某个数组插入到第 k 个位置，为了避免大规模的数据搬移，可以直接将第 k 位的数据搬移到数组元素的最后，把新的元素直接放入第 k 个位置。此时时间复杂度为O(1)

* 扩展

  * 为什么大多数编程语言中，数组要从 0 开始编号，而不是从 1 开始呢？

    从数组存储的内存模型上来看，“下标”最确切的定义应该是“偏移（offset）”.如果用 a 来表示数组的首地址，a[0] 就是偏移为 0 的位置，也就是首地址，a[k] 就表示偏移 k 个 type_size 的位置，所以计算 a[k] 的内存地址只需要用这个公式：

  <b>a[k]_address = base_address + k × type_size</b>

    但是，如果数组从 1 开始计数，那我们计算数组元素 a[k] 的内存地址就会变为：

  <b>a[k]_address = base_address + (k-1) × type_size</b>

    对比两个公式，我们不难发现，从 1 开始编号，每次随机访问数组元素都多了一次减法运算，对于 CPU 来说，就是多了一次减法指令。数组作为非常基础的数据结构，通过下标随机访问数组元素又是其非常基础的编程操作，效率的优化就要尽可能做到极致。所以为了减少一次减法操作，数组选择了从 0 开始编号，而不是从 1 开始。

## 链表(Linked list)

* 简述 空间上不一定连续 每个结点包括两个部分：一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域

* 时间复杂度
  * prepend O(1)
  * append O(1)
  * lookup O(n)
  * insert O(1)
  * delete O(1)

### 跳表(skip list)

有序链表加多级索引

对标平衡树(AVL Tree) 和二分查找,是一种 插入/删除/搜索 都是 O(logn)的数据结构

* 优点: 原理简单，容易实现，方便扩展，效率更高

运用: Redis LevelDB LRU

## 实战题目

### [盛最多水的容器](./container_with_most_water_test.go)

[leetcode](https://leetcode-cn.com/problems/container-with-most-water/)

### [移动零](./move_zeroes_test)

[leetcode](https://leetcode-cn.com/problems/container-with-most-water/)

### [移动零](./move_zeroes_test.go)

[leetcode](https://leetcode-cn.com/problems/move-zeroes/)

### [爬楼梯](./climbing_stairs.go)

[leetcode](https://leetcode-cn.com/problems/climbing-stairs/)

### [三数之和](./three_sum_test.go)

[leetcode](https://leetcode-cn.com/problems/3sum/)

### [数组形式的整数加法](./add_array_integer_test.go)

[leetcode](https://leetcode-cn.com/problems/add-to-array-form-of-integer/)

### [两数之和](./two_sum_test.go)

[leetcode](https://leetcode-cn.com/problems/two-sum/)
