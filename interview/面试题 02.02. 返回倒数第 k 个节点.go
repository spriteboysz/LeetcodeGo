/**
 * Author: Deean
 * Date: 2022-10-16 22:12
 * FileName: interview/面试题 02.02. 返回倒数第 k 个节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func kthToLast(head *lib.ListNode, k int) int {
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func main() {
	head := lib.Str2List("[1,2,3,4,5]")
	fmt.Println(kthToLast(head, 2))
}
