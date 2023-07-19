package main

import (
	"fmt"
)

/**
 * @Time    : 2023/7/19 09:31
 * @File    : slice2.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : GO语言面试宝典-数组和切片有什么异同？
 */

func main() {
	// 底层数组可以被多个slice同时指向
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	// s1 - [2, 3, 4]
	// 从s1的index-2开始，而s1的index-2指向的是slice的index-5，s1的index-6指向的是slice的index-9
	s2 := s1[2:6:7]
	// s2 - [4, 5, 6, 7]
	s2 = append(s2, 100)
	// s2 - [4, 5, 6, 7, 100]
	s2 = append(s2, 200)
	// s2 - [4, 5, 6, 7, 100, 200]
	s1[2] = 20
	// s1 - [2, 3, 20]
	// slice - [0, 1, 2, 3, 20, 5, 6, 7, 100, 9]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
