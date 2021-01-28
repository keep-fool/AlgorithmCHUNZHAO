# 学习笔记

## 目录

[作业](#实战题目)

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

* [盛最多水的容器](./container_with_most_water_test.go)

  [leetcode](https://leetcode-cn.com/problems/container-with-most-water/)

```go
// 解法一 暴力 双重遍历
func maxArea1(height []int) int {
  var max = 0
  for i := 0; i < len(height); i++ {
    for j := i + 1; j < len(height); j++ {
      min := func(a, b int) int {
        if a > b {
          return b
        }
        return a
      }(height[i], height[j])
      if max < min*(j-i) {
        max = min * (j - i)
      }
    }
  }
  return max
}
```

```go
// 解法二 双指针 夹逼
func maxArea2(height []int) int {
  o := 0
  i, j := 0, len(height)-1
  for i != j {
    hi, hj := height[i], height[j]
    s := (j - i) * func(a, b int) int {
      if a > b {
        return b
      }
      return a
    }(height[i], height[j])
    if s > o {
      o = s
    }
    if hi > hj {
      j--
    } else {
      i++
    }
  }
  return o
}
```

* [移动零](./move_zeroes_test.go)

  [leetcode](https://leetcode-cn.com/problems/move-zeroes/)

```go
// moveZeroes1 解法一 双指针
func moveZeroes1(nums []int) {
  if len(nums) == 0 {
    return
  }
  i, j := 0, 0
  for i < len(nums) {
    if nums[i] != 0 {
      nums[i], nums[j] = nums[j], nums[i]
      i++
      j++
    } else {
      i++
    }
  }
}
```

* [爬楼梯](./climbing_stairs_test.go)

  [leetcode](https://leetcode-cn.com/problems/climbing-stairs/)

```go
// 斐波那契
func climbStairs(n int) int {
  if n < 2 {
    return n
  }
  v1, v2, res := 1, 1, 2
  for i := 2; i <= n; i++ {
    res = v1 + v2
    v1, v2 = v2, res
  }
  return res
}
```

* [三数之和](./three_sum_test.go)

  [leetcode](https://leetcode-cn.com/problems/3sum/)

```go

// 暴力  理解题意
func threeSum1(nums []int) [][]int {
  sort.Ints(nums)
  fmt.Println(nums)
  var out [][]int
  var re = make(map[string]int)
  for i := 0; i < len(nums); i++ {
    if nums[i] > 0 {
      return out
    }
    if i > 0 && nums[i] == nums[i-1] {
      return out
    }
    for j := i + 1; j < len(nums); j++ {
      for m := j + 1; m < len(nums); m++ {
        if nums[i]+nums[j]+nums[m] == 0 {
          key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[m])
          if _, ok := re[key]; !ok {
            re[key] = 0
            out = append(out, []int{nums[i], nums[j], nums[m]})
          }
        }
      }
    }
  }
  return out
}

// 双指针
func threeSum3(nums []int) [][]int {
  sort.Ints(nums)
  var (
    out  [][]int
    lens = len(nums)
  )
  if lens < 3 {
    return out
  }
  for i := 0; i < len(nums); i++ {
    if nums[i] > 0 {
      return out
    }
    if i > 0 && nums[i] == nums[i-1] {
      continue
    }
    // 双指针
    var j, m = i + 1, lens - 1
    for j < m {
      sum := nums[i] + nums[j] + nums[m]
      switch {
      case sum == 0:
        out = append(out, []int{nums[i], nums[j], nums[m]})
        for j < m && nums[j] == nums[j+1] {
          j += 1
        }
        for j < m && nums[m] == nums[m-1] {
          m -= 1
        }
        j, m = j+1, m-1
      case sum > 0:
        m -= 1
      case sum < 0:
        j += 1
      }
    }
  }
  return out
}
```

* [数组形式的整数加法](./add_array_integer_test.go)

  [leetcode](https://leetcode-cn.com/problems/add-to-array-form-of-integer/)

```go
// 翻转
func addToArrayForm1(A []int, K int) []int {
  var (
    lens = len(A)
    add  int
  )
  A = reverse(A)
  for i := 0; i < lens || K != 0; i, K = i+1, K/10 {
    fmt.Println(i)
    if i >= lens {
      A = append(A, 0)
    }
    sum := K%10 + A[i] + add
    add = sum / 10
    A[i] = sum % 10
  }
  if add > 0 {
    A = append(A, add)
  }
  A = reverse(A)
  return A
}

func reverse(sli []int) []int {
  for i := 0; i < len(sli)/2; i++ {
    sli[i], sli[len(sli)-1-i] = sli[len(sli)-1-i], sli[i]
  }
  return sli
}
```

* [有效的字母异位词](./valid_anagram_test.go)

  [leetcode](https://leetcode-cn.com/problems/valid-anagram/)

```go
func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  ss := [26]int{}
  ts := [26]int{}
  for i := 0; i < len(s); i++ {
    index := s[i] - 'a'
    ss[index]++
  }
  fmt.Print(ss)
  for i := 0; i < len(t); i++ {
    index := t[i] - 'a'
    ts[index]++
  }
  fmt.Print(ss)
  return ss == ts
}
```

* [删除排序数组中的重复项](./remove_duplicates_test.go)

  [leetcode](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)

```go
// 快慢指针
func removeDuplicates(nums []int) int {
  lens := len(nums)
  var i, j = 0, 1
  for j < lens {
    if nums[i] == nums[j] {
      j++
    } else {
      i++
      nums[i] = nums[j]
    }
  }
  fmt.Println(nums)
  return i + 1
}
```

* [两数之和](./two_sum_test.go)

  [leetcode](https://leetcode-cn.com/problems/two-sum/)

```go
// 暴力解法
func twoSum(nums []int, target int) []int {
  l := len(nums)
  out := make([]int, 2)
  for i := 0; i < l; i++ {
    for j := i + 1; j < l; j++ {
      if nums[i]+nums[j] == target {
        out[0] = i
        out[1] = j
        return out
      }
    }
  }
  return out
}

// 哈希
func twoSum1(nums []int, target int) []int {
  result := []int{}
  m := make(map[int]int)
  for i, k := range nums {
    if value, exist := m[target-k]; exist {
      result = append(result, value)
      result = append(result, i)
    }
    m[k] = i
  }
  return result
}
```

* [66. 加一](./plus_one_test.go)

  [leetcode](https://leetcode-cn.com/problems/plus-one/)

```go
func plusOne(digits []int) []int {
  for i := len(digits) - 1; i >= 0; i-- {
    digits[i] += 1
    digits[i] %= 10
    if digits[i] != 0 {
      return digits
    }
  }
  return append([]int{1}, digits...)
}
```

* [189. 旋转数组](./rotate_array_test.go)

  [leetcode](https://leetcode-cn.com/problems/rotate-array/)

```go
func rotate(nums []int, k int) []int {
  res := []int{}
  res = append(nums[k:], nums[:k]...)
  nums = []int{}
  nums = res
  return nums
}

func rotate1(nums []int, k int) {
  reverse(nums)
  reverse(nums[:k%len(nums)])
  reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
  for i := 0; i < len(arr)/2; i++ {
    arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
  }
}
```

* [412. Fizz Buzz](./fizz_buzz_test.go)

  [leetcode](https://leetcode-cn.com/problems/fizz-buzz/)

```go
// 常规暴力做法
func fizzBuzz(n int) []string {
  out := []string{}
  for i := 1; i <= n; i++ {
    switch {
    case i%15 == 0:
      out = append(out, "FizzBuzz")
    case i%3 == 0:
      out = append(out, "Fizz")
    case i%5 == 0:
      out = append(out, "Buzz")
    default:
      out = append(out, fmt.Sprint(i))
    }
  }
  return out
}

func fizzBuzz1(n int) []string {
  out := []string{}
  for i := 1; i <= n; i++ {
    str := ""
    if i%3 == 0 {
      str += "Fizz"
    }
    if i%5 == 0 {
      str += "Buzz"
    }
    if str == "" {
      str += fmt.Sprint(i)
    }
    out = append(out, str)
  }
  return out
}

var vals = []string{
  "%d", "%d", "Fizz%.T", "%d", "Buzz%.T", "Fizz%.T", "%d", "%d",
  "Fizz%.T", "Buzz%.T", "%d", "Fizz%.T", "%d", "%d", "FizzBuzz%.T",
}

func fizzBuzz2(n int) []string {
  ret := make([]string, n)
  for i := range ret {
    ret[i] = fmt.Sprintf(vals[i%len(vals)], i+1)
  }
  return ret
}
```

* [21. 合并两个有序链表](./merge_two_sorted_lists_test.go)

  [leetcode](https://leetcode-cn.com/problems/merge-two-sorted-lists/solution/)
