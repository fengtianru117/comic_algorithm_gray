package main

import "fmt"

/*
 * https://blog.csdn.net/youngwhz1/article/details/83026263
 *
 * 在Slice指定位置插入元素。
 * params:
 *   s: slice对象，类型为[]interface{}
 *   index: 要插入元素的位置索引
 *   value: 要插入的元素
 * return:
 *   已经插入元素的slice，类型为[]interface{}
 */

func SliceInsert(s []interface{}, index int, value interface{}) []interface{} {
	// 数组的底层数据是不变的 这么写会导致共享部分被篡改
	// rear := s[index:]
	rear := append([]interface{}{}, s[index:]...)
	fmt.Println(rear...)
	s = append(s[:index], value)
	fmt.Println(s)
	s = append(s, rear...)
	return s
}

func main() {
	var a = []interface{}{1, 2, 3, 4}
	// 数组的底层数据是不变的
	var b = []interface{}{1, 2, 3, 4}
	temp := a[2:]
	fmt.Println(temp)
	b[2] = 10
	fmt.Println(temp)
	fmt.Println("#########################")
	fmt.Println(SliceInsert(a, 2, 10))
}
