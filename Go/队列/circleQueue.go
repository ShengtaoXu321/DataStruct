package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	MaxSize int    // 4
	array   [4]int // 数组
	head    int    // 指向队列的队首
	tail    int    // 指向队列的队尾
}

func main() {
	// 先创建一个队列
	cqueue := CircleQueue{
		MaxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	// 添加数据
	for true {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出队列")
		fmt.Println()
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数")
			fmt.Scanln(&val)
			err := cqueue.Push(val) // 调用一个方法了
			if err != nil {
				fmt.Println("加入队列失败！", err)
			}
			fmt.Println("加入队列成功！")
		case "get":
			val, err := cqueue.Pop()
			if err != nil {
				fmt.Println("队列取出数据失败！", err)
			} else {
				fmt.Println("从队列中取出的数据是：", val)
			}
		case "show":
			cqueue.ShowQueue()
		case "exit":
			os.Exit(0)

		}
	}
}

// 入队列方法：-- AddQueue(也叫做push)     GetQueue(也叫做pop)
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() == true {
		return errors.New("queue is full!")
	}
	// 分析出this.tail 在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val // 把值赋给尾部
	//this.tail++   错误用法
	this.tail = (this.tail + 1) % this.MaxSize
	return
}

// 出队列 pop()方法
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() == true {
		return -1, errors.New("queue is empty!")
	}
	// 取出,head是指向队首，并且含队首元素
	val = this.array[this.head]
	this.tail = (this.tail + 1) % this.MaxSize
	return val, err
}

// 显示队列
func (this *CircleQueue) ShowQueue() {
	fmt.Println("环形队列情况如下。。。。")
	// 取出当前队列有多少个元素
	size := this.LenQueue() // 获取有多少个元素
	if size == 0 {
		fmt.Println("队列为空！")
	}
	// 设计一个临时变量，指向head
	temp := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d", temp, this.array[temp])
		temp = (temp + 1) % this.MaxSize
	}
	fmt.Println()
}

// 判断环形队列为满 的方法
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.MaxSize == this.head
}

// 判断环形队列为空
func (this *CircleQueue) IsEmpty() bool {
	return (this.tail) == (this.head)
}

// 取出环形队列有多少个元素
func (this *CircleQueue) LenQueue() int {
	// 这是一个关键的算法
	return (this.tail + this.MaxSize - this.head) % this.MaxSize
}
