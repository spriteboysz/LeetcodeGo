/**
 * Author: Deean
 * Date: 2022/11/25 23:15
 * FileName: interview/面试题 02.06. 回文链表.go
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
	for i := 0; i < n/2; i++ {
		if values[i] != values[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(lib.Str2List("[1,2,2,1]")))
}
