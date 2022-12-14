/**
 * Author: Deean
 * Date: 2022/12/13 23:37
 * FileName: interview/面试题 02.08. 环路检测.go
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
