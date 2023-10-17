/**
 * Author: Deean
 * Date: 2023-10-15 23:35
 * FileName: LCP/LCR 123. 图书整理 I.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reverseBookList(head *lib.ListNode) []int {
	values := []int{}
	if head == nil {
		return values
	}
	for head != nil {
		values = append([]int{head.Val}, values...)
		head = head.Next
	}
	return values
}

func main() {
	fmt.Println(reverseBookList(lib.Str2List("[3,6,4,1]")))
}
