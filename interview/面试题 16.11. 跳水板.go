/**
 * Author: Deean
 * Date: 2022/11/25 23:09
 * FileName: interview/面试题 16.11. 跳水板.go
 * Description:
 */

package main

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	boards := []int{}
	for i := 0; i <= k; i++ {
		boards = append(boards, shorter*(k-i)+longer*i)
	}
	return boards
}

func main() {
	fmt.Println(divingBoard(1, 2, 3))
}
