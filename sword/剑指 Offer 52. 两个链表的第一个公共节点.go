/**
 * Author: Deean
 * Date: 2022-10-21 23:04
 * FileName: sword/剑指 Offer 52. 两个链表的第一个公共节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func getIntersectionNode2(headA, headB *lib.ListNode) *lib.ListNode {
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
	fmt.Println(getIntersectionNode2(headA, headB))
}
