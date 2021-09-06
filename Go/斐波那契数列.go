package main

import "fmt"

// 尾递归实现斐波那契函数
// 斐波那契数列： 1 1 2 3 5 8 ... N-1 N 2N-1

func main() {
	fmt.Println(5, 1,1)
}

func Fona(n, a,b int) int {
	if n ==0 {
		return a
	}
	return Fona(n-1, b, a+b)
}

