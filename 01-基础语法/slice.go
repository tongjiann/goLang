package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 1) // 追加1个元素
	fmt.Println(arr)
	arr = append(arr, 1, 2, 3) // 追加多个元素, 手写解包方式
	fmt.Println(arr)
	arr = append(arr, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(arr)
	arr = append([]int{0}, arr...) // 在开头位置添加1个元素
	fmt.Println(arr)
	arr = append([]int{-3, -2, -1}, arr...) // 在开头添加1个切片
	fmt.Println(arr)
	for i := 0; i < 5; i++ {
		copy(arr[i+1:], arr[i:])
	}
	fmt.Println(arr)
}
