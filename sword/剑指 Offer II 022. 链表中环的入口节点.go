/**
 * Author: Deean
 * Date: 2022/12/12 23:33
 * FileName: sword/剑指 Offer II 022. 链表中环的入口节点.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func detectCycle(head *lib.ListNode) *lib.ListNode {
	nodes := map[*lib.ListNode]bool{}
	for head != nil {
		if nodes[head] {
			return head
		}
		nodes[head] = true
		head = head.Next
	}
	return nil
}

func main() {
	fmt.Println(detectCycle(lib.Str2List("[3,2,0,-4]")))
}
