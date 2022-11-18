/**
 * Author: Deean
 * Date: 2022/11/18 23:06
 * FileName: sword/剑指 Offer 25. 合并两个排序的链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func mergeTwoLists(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}

func main() {
	fmt.Println(lib.List2String(mergeTwoLists(lib.Str2List("[1,2,3]"), lib.Str2List("[1,3,4]"))))
}
