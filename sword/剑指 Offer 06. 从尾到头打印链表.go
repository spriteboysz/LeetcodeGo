/**
 * Author: Deean
 * Date: 2022/11/13 23:58
 * FileName: sword/剑指 Offer 06. 从尾到头打印链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reversePrint(head *lib.ListNode) []int {
	if head == nil {
		return []int{}
	}
	values := []int{}
	for head != nil {
		values = append([]int{head.Val}, values...)
		head = head.Next
	}
	return values
}

func main() {
	fmt.Println(reversePrint(lib.Str2List("[1,2,3]")))
}
