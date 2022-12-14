/**
 * Author: Deean
 * Date: 2022/12/14 22:03
 * FileName: algorithm/P0203. 移除链表元素.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func removeElements(head *lib.ListNode, val int) *lib.ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func main() {
	fmt.Println(lib.List2String(removeElements(lib.Str2List("[1,2,6,3,4,5,6]"), 6)))
}
