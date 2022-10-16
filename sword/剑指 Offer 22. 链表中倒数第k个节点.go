/**
 * Author: Deean
 * Date: 2022-10-16 16:52
 * FileName: sword/剑指 Offer 22. 链表中倒数第k个节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getKthFromEnd(head *lib.ListNode, k int) *lib.ListNode {
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	head := lib.Str2List("[1,2,3,4,5]")
	fmt.Println(lib.List2String(getKthFromEnd(head, 2)))
}
