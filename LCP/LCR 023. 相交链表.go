/**
 * Author: Deean
 * Date: 2023-10-14 22:46
 * FileName: LCP/LCR 023. 相交链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getIntersectionNode(headA, headB *lib.ListNode) *lib.ListNode {
	visited := map[*lib.ListNode]bool{}
	for headA != nil {
		visited[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if visited[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	fmt.Println(lib.List2String(getIntersectionNode(lib.Str2List("[4,1,8,4,5]"), lib.Str2List("[4,1,8,4,5]"))))
}
