/**
 * Author: Deean
 * Date: 2023-10-19 22:49
 * FileName: LCR/LCR 171. 训练计划 V.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getIntersectionNode5(headA, headB *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	hash := map[*ListNode]bool{}
	for headA != nil {
		hash[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if hash[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	headA, headB := lib.Str2List("[4,1,8,4,5]"), lib.Str2List("[4,1,8,4,5]")
	fmt.Println(lib.List2String(getIntersectionNode5(headA, headB)))
}
