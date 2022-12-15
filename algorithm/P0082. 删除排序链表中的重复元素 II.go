/**
 * Author: Deean
 * Date: 2022/12/14 22:44
 * FileName: algorithm/P0082. 删除排序链表中的重复元素 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func deleteDuplicatesII(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return head
	}
	dummy := &lib.ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	fmt.Println(lib.List2String(deleteDuplicatesII(lib.Str2List("[1,1,1,2,3]"))))
}
