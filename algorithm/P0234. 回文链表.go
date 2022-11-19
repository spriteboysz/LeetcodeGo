/**
 * Author: Deean
 * Date: 2022/11/19 16:20
 * FileName: algorithm/P0234. 回文链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func isPalindrome(head *lib.ListNode) bool {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	n := len(values)
	for i := 0; i < len(values); i++ {
		if values[i] != values[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(lib.Str2List("[1,2,2,1]")))
}
