/**
 * Author: Deean
 * Date: 2022/11/19 16:13
 * FileName: sword/剑指 Offer II 027. 回文链表.go
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
	fmt.Println(isPalindrome(lib.Str2List("[1,2,3,3,2,1]")))
}
