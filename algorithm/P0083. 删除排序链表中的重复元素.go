/**
 * Author: Deean
 * Date: 2022/12/4 22:03
 * FileName: algorithm/P0083. 删除排序链表中的重复元素.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deleteDuplicates(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	fmt.Println(lib.List2String(deleteDuplicates(lib.Str2List("[1,1,2,3,3]"))))
}
