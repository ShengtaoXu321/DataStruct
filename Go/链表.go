package main

import "fmt"

/*
什么是链表：
由一个个数据节点组成，是一个递归结构，要么为空，要么指向另外一个数据节点的引用
*/

// 最简单的链表

// 结构体定义链表
type LinkNode struct {
	Data int64
	NextNode *LinkNode
}
// 链表的特点：结构体中有两个字段：一个是存放数据，一个是指向下一个节点的指针
/* 链表的划分：单链表、双链表、循环单链表、循环双链表
单链表：单向的，只有一个方向，不能往回走
双链表：双向的，既可以找到之前的节点，也可以找到后面的
循环链表：就是一直往下找数据节点，最后回到自己那个节点，形成一个回路。  循环单链表和循环双链表的区别：一个只能一个方向走，一个两个方向都可以走
*/

/*
问题：make和new的区别
1. make和new都是用来分配内存的内建函数，且在堆上分配。make既会分配内存，也会初始化，而new只是将内存清零，并没有初始化
2. make的返回的是引用类型本身，而new返回的是指向类型的指针
3. make只能用来分配和初始化类型为slice, map, channel的数据，new可以分配任意类型的数据
*/

func main() {
	// 新的节点1
	node:=new(LinkNode)
	node.Data = 2

	// 新的节点2
	node1:=new(LinkNode)
	node1.Data=3
	node.NextNode=node1  // node1 链接到 node 节点上

	// 新的节点3
	node2:=&LinkNode{
		Data: 4,
	}
	node1.NextNode=node2

	// 按顺序打印数据 // node2 链接到 node1 节点上

	//fmt.Println(node)
	//nowNode:=node
	//fmt.Println(nowNode)
	//for  {
	//	if nowNode!=nil{
	//		// 打印节点值
	//		fmt.Println(nowNode.Data)
	//		// 获取下一个节点
	//		nowNode=nowNode.NextNode
	//		continue
	//	}
	//	// 如果下一个节点为空
	//	break
	//}
	for  {
		if node!=nil{
			fmt.Println(node.Data)
			node=node.NextNode
			continue
		}
		break
	}



}