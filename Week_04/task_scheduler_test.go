package week04

import (
	"fmt"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	var (
		tasks = "AAABBB"
		n     = 2
	)
	t.Log(leastInterval([]byte(tasks), n))
}

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
