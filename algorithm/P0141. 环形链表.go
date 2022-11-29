/**
 * Author: Deean
 * Date: 2022/11/28 22:30
 * FileName: algorithm/P0141. 环形链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func hasCycle(head *lib.ListNode) bool {
	hash := map[*lib.ListNode]bool{}
	for head != nil {
		if hash[head] {
			return true
		}
		hash[head] = true
		head = head.Next
	}
	return false
}

func main() {
	fmt.Println(hasCycle(lib.Str2List("[3,2,0,-4]")))
}
