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

* [860. 柠檬水找零](./lemonade_change_test.go)

    [leetcode](https://leetcode-cn.com/problems/lemonade-change/)

```go
func lemonadeChange(bills []int) bool {
    five, ten := 0, 0
    for _, bill := range bills {
        if bill == 5 {
            five++
        } else if bill == 10 {
            if five == 0 {
                return false
            }
            five--
            ten++
        } else {
            if five > 0 && ten > 0 {
                five--
                ten--
            } else if five >= 3 {
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}
```

* [529. 扫雷游戏](./minesweeper_test.go)

    [leetcode](https://leetcode-cn.com/problems/minesweeper/)

```go
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
    x, y := click[0], click[1]
    if board[x][y] == 'M' {
        board[x][y] = 'X'
    } else {
        dfs(board, x, y)
    }
    return board
}

func dfs(board [][]byte, x, y int) {
    cnt := 0
    for i := 0; i < 8; i++ {
        tx, ty := x+dirX[i], y+dirY[i]
        if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
            continue
        }
        if board[tx][ty] == 'M' {
            cnt++
        }
    }
    if cnt > 0 {
        board[x][y] = byte(cnt + '0')
    } else {
        board[x][y] = 'B'
        for i := 0; i < 8; i++ {
            tx, ty := x+dirX[i], y+dirY[i]
            if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
                continue
            }
            dfs(board, tx, ty)
        }
    }
}
```

* [51. N 皇后](./binary_tree_traversal_test.go)

    [leetcode](https://leetcode-cn.com/problems/n-queens/)

```go
func solveNQueens(n int) [][]string {
    solution := make([][]string, 0)
    cUsed, dUsed := make([]bool, n), make([]bool, (n*2-1)*2)
    fmt.Println(cUsed)
    fmt.Println(dUsed)
    backtrack(&solution, buildNxNBoard(n), cUsed, dUsed, n, 0)
    return solution
}

func backtrack(solution *[][]string, board [][]byte, cUsed, dUsed []bool, n, y int) {
    if y == n {
        boardCopy := make([]string, n)
        for i := range boardCopy {
            boardCopy[i] = string(board[i])
        }
        *solution = append(*solution, boardCopy)
    } else {
        for x := 0; x < n; x++ {
            if !cUsed[x] && !dUsed[y-x+n-1] && !dUsed[y+x+n*2-1] {
                cUsed[x], dUsed[y-x+n-1], dUsed[y+x+n*2-1] = true, true, true
                board[y][x] = 'Q'
                backtrack(solution, board, cUsed, dUsed, n, y+1)
                board[y][x] = '.'
                cUsed[x], dUsed[y-x+n-1], dUsed[y+x+n*2-1] = false, false, false
            }
        }
    }
}

func buildNxNBoard(n int) [][]byte {
    result := make([][]byte, n)
    str := make([]byte, n)
    for i := range str {
        str[i] = '.'
    }
    for i := range result {
        result[i] = append([]byte{}, str...)
    }
    return result
}
```
