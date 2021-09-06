package main

import "fmt"

type ValNode struct {
	Row int
	Col int
	Val int
}

func main() {
	// 1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 表示黑子
	chessMap[2][3] = 2 // 表示蓝子

	// 2. 输出原始数组看看
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 思路
	// (1) 遍历chessMap， 如果我们发现有一个元素的值不为0，创建一个node结构体
	// (2) 将其放入到对应的切片即可
	var sparseArr []ValNode // 是一个ValNode类型的数组

	// 标准的稀疏数组，除了标注非0值，还需要记录元素的二维数组的规模（行和列，默认值）
	valnode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArr = append(sparseArr, valnode)

	for i, v := range chessMap {
		for j, v2 := range v {
			// 判断是否为0
			if v2 != 0 {
				// 创建一个ValNode 值结点
				valnode := ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				// 将 值结点数据追加到创建的切片中
				sparseArr = append(sparseArr, valnode)
			}
		}
	}
	fmt.Println("稀疏数组整体打印是：", sparseArr)

	// 循环遍历输出稀疏数组
	fmt.Println("当前的稀疏数组循环输出打印是")
	for i, valnode := range sparseArr {
		fmt.Printf("第%d个值: %d %d %d\n", i, valnode.Row, valnode.Col, valnode.Val)
	}
	// 4. 将这个稀疏数组，存盘 d:/chessmap.data

	// 5. 如何恢复这个原始数组 -- 续上盘
	// 思路：
	// 1. 打开这个d:/chessmap.data ==》 恢复原始数组
	// 2. 这里使用稀疏数组恢复

	var chessMap2 [11][11]int // 这里使用的是数组，还可以使用切片

	for i, valnode := range sparseArr {
		if i != 0 { // 跳过第一行的值
			chessMap2[valnode.Row][valnode.Col] = valnode.Val
		}
	}
	fmt.Println(chessMap2)

}
