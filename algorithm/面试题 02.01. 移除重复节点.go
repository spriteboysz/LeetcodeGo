/**
 * Author: Deean
 * Date: 2022/11/15 23:19
 * FileName: algorithm/面试题 02.01. 移除重复节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func removeDuplicateNodes(head *lib.ListNode) *lib.ListNode {
	ob := head
	for ob != nil {
		oc := ob
		for oc.Next != nil {
			if oc.Next.Val == ob.Val {
				oc.Next.Next = oc.Next
			} else {
				oc = oc.Next
			}
		}
		ob = ob.Next
	}
	return head
}

func main() {
	fmt.Println(lib.List2String(removeDuplicateNodes(lib.Str2List("[1, 1, 1, 1, 2]"))))
}
