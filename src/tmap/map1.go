package main

import "fmt"

/**
 * @Time    : 2023/7/13 17:16
 * @File    : map1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 探讨 map（映射）
 */

func main() {
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	val1, flag1 := map1["one"]
	fmt.Println(val1, flag1)
	val2 := map1["two"]
	if val2 != 0 {
		fmt.Println(val2, "True")
	}
	val3 := map1["three"]
	if val3 != 0 {
		fmt.Println(val3, "True")
	} else {
		fmt.Println(val3, "False")
	}

	delete(map1, "two")

	for key, value := range map1 {
		fmt.Println(key, value)
	}
	//map2 := map[string]string{}
	//map2["two"] = "two"
	//fmt.Println(map2)

	//var map3 map[string]string
	//map3["three"] = "three"
	//fmt.Println(map3)
	// 爱奇艺 140
	// B站 + QQ 音乐 288
	// 淘宝 88
}
