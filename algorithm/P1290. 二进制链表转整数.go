/**
 * Author: Deean
 * Date: 2022-10-16 16:11
 * FileName: algorithm/P1290. 二进制链表转整数.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getDecimalValue(head *lib.ListNode) int {
	num := 0
	for head != nil {
		num = num*2 + head.Val
		head = head.Next
	}
	return num
}

func main() {
	head := lib.Str2List("[1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]")
	fmt.Println(getDecimalValue(head))
}
