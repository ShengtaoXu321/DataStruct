package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	MaxSize int
	Array   [5]int // 数组 ==》 模拟队列
	front   int    // 表示指向队首
	rear    int    // 表示指向队列的尾部
}

func main() {
	// 先创建一个队列
	queue := Queue{
		MaxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val) // 调用一个方法了
			if err != nil {
				fmt.Println("加入队列失败！", err)
			}
			fmt.Println("加入队列成功！")
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("队列取出数据失败！", err)
			} else {
				fmt.Println("从队列中取出的数据是：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)

		}
	}
}

// 添加数据到队列 --这不是一个函数，是方法了
func (this *Queue) AddQueue(val int) (err error) {
	// 第一步：先判断队列是否已满
	if this.rear == this.MaxSize-1 { // 重要：rear是队列的尾部（含最后的元素）
		return errors.New("queue full!")
	}
	// 队列未满，添加数据时，rear后移
	this.rear++
	// 添加数据
	this.Array[this.rear] = val
	return
}

// 显示队列,找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是。。。")
	// this.front不包含队首元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d", i, this.Array[i])
	}
	fmt.Println()
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.front == this.rear { // 队列空
		return 0, errors.New("queue null!")
	}
	// 不为空， 队首前进
	this.front++
	// 取出值
	val = this.Array[this.front]
	return val, err
}
