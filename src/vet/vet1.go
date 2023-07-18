package vet

import "fmt"

/**
 * @Time    : 2023/7/11 16:19
 * @File    : vet1.go
 * @Project : chapter3
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 尝试 go vet
 */

//(base) ☁  vet  go vet vet1.go
//# command-line-arguments
//./vet1.go:15:12: fmt.Printf call has arguments but no formatting directives

func main() {
	fmt.Printf("Pi is:", 3.14)
}

// 7.5*8/100 = 0.6
// 6.5*8/100 = 0.52
