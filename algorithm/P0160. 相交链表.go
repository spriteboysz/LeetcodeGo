/**
 * Author: Deean
 * Date: 2022-10-21 23:45
 * FileName: algorithm/P0160. 相交链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getIntersectionNode(headA, headB *lib.ListNode) *lib.ListNode {
	hash := map[*lib.ListNode]bool{}
	cur := headA
	for cur != nil {
		hash[cur] = true
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		if hash[cur] {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func main() {
	headA := lib.Str2List("[1,2,3,4]")
	headB := lib.Str2List("[1,2,3,4,5]")
	fmt.Println(getIntersectionNode(headA, headB))
}
