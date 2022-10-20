/**
 * Author: Deean
 * Date: 2022-10-20 23:51
 * FileName: sword/剑指 Offer II 023. 两个链表的第一个重合节点.go
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
	headA := lib.Str2List("[0,9,1,2,4]")
	headB := lib.Str2List("[2,4]")
	fmt.Println(getIntersectionNode(headA, headB))
}
