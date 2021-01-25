package week01

import (
	"fmt"
	"testing"
)

// 写一个程序，输出从 1 到 n 数字的字符串表示。

// 1. 如果 n 是3的倍数，输出“Fizz”；

// 2. 如果 n 是5的倍数，输出“Buzz”；

// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

func TestFizzBuzz(t *testing.T) {
	n := 15
	t.Log(fizzBuzz(n))
	t.Log(fizzBuzz1(n))
}

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
