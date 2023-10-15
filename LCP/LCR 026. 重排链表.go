/**
 * Author: Deean
 * Date: 2023-10-14 23:52
 * FileName: LCP/LCR 026. 重排链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reorderList(head *lib.ListNode) {
	type ListNode = lib.ListNode
	if head == nil {
		return
	}
	nodes1, nodes2, cur := []*ListNode{}, []*ListNode{}, head
	for cur != nil {
		nodes1 = append(nodes1, cur)
		cur = cur.Next
	}
	n := len(nodes1)
	for i := 0; i <= n/2; i++ {
		nodes2 = append(nodes2, nodes1[i])
		nodes2 = append(nodes2, nodes1[n-1-i])
	}
	for i := 0; i < n-1; i++ {
		nodes2[i].Next = nodes2[i+1]
	}
	nodes2[n-1].Next = nil
	fmt.Println(lib.List2String(head))
}
func main() {
	reorderList(lib.Str2List("[1,2,3,4]"))
	reorderList(lib.Str2List("[1,2,3,4,5]"))
}
