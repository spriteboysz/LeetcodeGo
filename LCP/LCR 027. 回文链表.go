/**
 * Author: Deean
 * Date: 2023-10-15 09:33
 * FileName: LCP/LCR 027. 回文链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isPalindrome2(head *lib.ListNode) bool {
	nodes := []int{}
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	for i, n := 0, len(nodes); i < n/2; i++ {
		if nodes[i] != nodes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome2(lib.Str2List("[1,2,3,3,2,1]")))
	fmt.Println("[1,2,3,4,5]")
}
