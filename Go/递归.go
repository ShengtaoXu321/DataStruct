package main

import "fmt"

// 递归实现阶乘
func main() {
	var n int
	fmt.Scan(&n)
	v:=Rescuvie(n)
	fmt.Println(v)
}

func Rescuvie(n int) (value int){
	if n ==0 {
		return 1
	}
	value=n*Rescuvie(n-1)
	return value
}

// 上述的递归方法使用了 运算符， 重复的调用都使得运算的链条不断加长，系统不得不使用栈进行数据保存和恢复
// 如果每次递归都要对越来越长的链进行运算，速度很慢，可能会导致栈溢出，程序崩溃等

// 使用 尾递归
func RescuvieTail(n,a int) int {   // a代表是尾数，a表示阶乘后的结果
	if n == 1{
		return a
	}

	return RescuvieTail(n-1, a)
	//RescuvieTail(5, 1)
	//RescuvieTail(4, 1*5)=RescuvieTail(4, 5)

}
// 尾部递归：递归函数在调用自身后直接传回其值，而不在对其再加运算
// 尾递归函数：如果一个函数中所有递归形式的调用都出现在函数的末尾