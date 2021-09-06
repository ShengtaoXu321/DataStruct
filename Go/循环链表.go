package main

import "fmt"

// 定义一个循环链表
type Ring struct {
	next, prev *Ring   // 前驱和后驱节点
	Value interface{}  // 数据
}

// 初始化一个空的循环链表 -- 封装一个方法
func (r *Ring) init() *Ring {
	r.next=r
	r.prev=r
	return r
}
// 绑定前驱和后驱节点都是自己，没有循环，时间复杂度为O(1)

// 创建一个指定大小N的循环链表，值全为空
func New(n int) *Ring {
	if n<=0{
		return nil
	}

	r:=new(Ring)
	p:=r
	for i:=0;i<n;i++{
		p.next=&Ring{prev: p}
		p=p.next
	}
	p.next=r
	r.prev=p
	return r
}  // 连续绑定前驱和后驱节点，时间复杂度O(n)


// 方法：获取上一个下一个节点
func (r * Ring) Next() *Ring {
	if r.next==nil{
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.prev==nil{
		return r.init()
	}
	return r.prev
}

// 获取第n个节点：因为链表是循环的，当n为负数，表示从前面往前遍历，否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next==nil{
		return r.init()
	}
	switch {
	case n<0:
		for ; n<0;n++{
			r=r.prev
		}
	case n>0:
		for ;n>0;n--{
			r=r.next
		}
	}
	return r
}

// 添加节点
func (r *Ring) Link(s *Ring) *Ring {
	n:=r.Next()   // 获取下一个节点
	if s!=nil{
		p:=s.Prev()   // 获取上一个节点
		r.next=s
		s.prev=r
		n.prev=p
		p.next=n
	}
	return n
}
/*添加一个新的节点：
插入的节点s是插入r后的一个新的节点，需要r的后驱节点n指向s的前驱结点p，s指向r  闭环
*/

// 用数组来实现链表
/*
数组和链表的比较
1. 数组是占用连续的空间，优点是占用空间小，查询快  缺点是移动和删除数据需要大量的空间
2. 链表的优点：移动和删除数据元素速度快，缺点是占用空间大，查找需要遍历
*/


func MakeArrLink()  {
	type ArrayLink struct {
		Data string
		NextIndex int64
	}

	var array [5]ArrayLink
	array[0]=ArrayLink{
		Data:      "I", // 值为I
		NextIndex: 3,   // 下一个节点的下标为3
	}
	array[1]=ArrayLink{"you",5}
	array[2]=ArrayLink{"like",4}
	array[3]=ArrayLink{"!",-1}   // -1表示没有下一个节点
	node:=array[0]
	// 输出
	for true {
		fmt.Println(node.Data)
		if node.NextIndex==-1{
			break
		}
		node=array[node.NextIndex]  // 下标需要移动
	}
}





