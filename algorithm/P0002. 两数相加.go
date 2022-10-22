/**
 * Author: Deean
 * Date: 2022-10-22 16:08
 * FileName: algorithm/P0002. 两数相加.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	dummy := &lib.ListNode{Val: -1}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &lib.ListNode{Val: carry % 10}
		cur = cur.Next
		carry /= 10
	}
	return dummy.Next
}

func main() {
	l1 := lib.Str2List("[9,9,9,9,9,9,9]")
	l2 := lib.Str2List("[9,9,9,9]")
	fmt.Println(lib.List2String(addTwoNumbers(l1, l2)))
}
