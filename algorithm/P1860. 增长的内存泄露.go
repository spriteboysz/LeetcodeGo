/**
 * Author: Deean
 * Date: 2022-11-06 23:03
 * FileName: algorithm/P1860. 增长的内存泄露.go
 * Description:
 */

package main

import "fmt"

func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for ; i <= memory1 || i <= memory2; i++ {
		if memory1 >= memory2 {
			memory1 -= i
		} else {
			memory2 -= i
		}
	}
	return []int{i, memory1, memory2}
}

func main() {
	fmt.Println(memLeak(8, 11))
}
