package main

import "fmt"

// 问题描述： 有A B C 三根杆子， A杆存放64个金盘，目标：把A杆上的金盘全部移动到C杆上，保持原有的顺序
/* 思想：
1. 借助C杆，将A杆前面的N-1个盘子，全部移动到B杆；再将A的最后一个盘子，直接移动到C杆。这时候A杆空
2. 剩下B杆的N-1个盘子，跟上述方法一致，借助A，将B的盘子，移动到C
*/

// 求次数
var total int

func main()  {
	//n:=4  // 64个盘子
	var n int
	fmt.Scan(&n)
	a:="a"// A杆
	b:="b"// B杆
	c:="c"// C杆

	tower(n,a,b,c)
	// 当 n=1 时，移动次数为 1
	// 当 n=2 时，移动次数为 3
	// 当 n=3 时，移动次数为 7
	// 当 n=4 时，移动次数为 15
	fmt.Println(total)

}

// 表示将N个盘子，从a杆，借助b杆移动到c杆
func tower(n int, a,b,c string)  {
	if n==1{
		//total=total+1
		total += 1
		fmt.Println(a,"->",c)
		return
	}

	// 递归
	tower(n-1,a,c,b)
	total += 1
	fmt.Println(a,"->",c)
	tower(n-1,b,a,c)
}