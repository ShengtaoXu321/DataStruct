package main

import "fmt"

// 二分查找 1 5 9 15 81 89 123 189 333，找出数字189
// 思路：先拿排好序的中位数与目标数字189对比，如果刚好匹配就结束。
// 如果中位数比目标数字大，就从左边找
// 如果中位数比目标数字小，就从右边找

// 递归

func main() {
	arr:=[]int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	tar:=189
	res:=BinarySearch(arr,tar,0,len(arr)-1)   // 减1是因为下标是从0开始
	fmt.Println(res)
}

func BinarySearch(array []int, target, first, end int) int {
	if first>end{   // 前提：排好序
		return  -1  // 起始小标大于末位置下标，出界
	}
	mid:=(first+end)/2  // 获取中间下标
	middNum:=array[mid] // 获取中间值

	// 进行判断
	if middNum == target{
		return mid  // 找到下标
	}else if middNum > target{
		// 中间数比目标数大，往左找
		return BinarySearch(array,target,0,mid-1)
	}else{
		// 中间数比目标数小，往右找
		return BinarySearch(array, target, mid+1, end)
	}

}