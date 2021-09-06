package main

import "sync"

/*
切片的底层
type slice struct{
	array unsafe.Pointer
	len	  int
	cap   int
}
涵盖： 1， 指向底层数组的指针	2. 切片的真正长度（实际元素占用的大小）		3. 切片的容量（底层固定数组的长度）
注意：
1. 扩容是针对cap进行扩容的
2. 切片的append操作必须要原来的变量去接，也就是 s1=append(s1,"1") 而不是_=append(s1,"1")
3. 切片的扩容策略：小于1024，扩容翻倍；大于1024，每次增加25%
*/

// 构造一个切片
type Slice struct {
	array []int  // 固定大小的数组，用满容量和满大小的切片来替代
	len int    // 真正的长度
	cap int    // 底层固定数组的长度
	lock sync.Locker  // 为了并发安全使用的锁
}

// 初始化数组