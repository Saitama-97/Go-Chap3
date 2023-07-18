package main

import "fmt"

/**
 * @Time    : 2023/7/13 09:38
 * @File    : slice1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 探讨 len (长度)和 cap (容量)
 */

func main() {
	slice1 := make([]int, 3, 5)
	//for i := 0; i < 3; i++ {
	//	fmt.Println(len(append(slice1, i+1)), cap(append(slice1, i+1)))
	//}
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(len(slice1), cap(slice1))
}
