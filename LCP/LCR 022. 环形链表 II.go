/**
 * Author: Deean
 * Date: 2023-10-14 22:42
 * FileName: LCP/LCR 022. 环形链表 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func detectCycle(head *lib.ListNode) *lib.ListNode {
	visited := map[*lib.ListNode]bool{}
	for head != nil {
		if visited[head] {
			return head
		}
		visited[head] = true
		head = head.Next
	}
	return nil
}

func main() {
	fmt.Println(lib.List2String(detectCycle(lib.Str2List("[3,2,0,-4]"))))
}
