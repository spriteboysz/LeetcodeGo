/**
 * Author: Deean
 * Date: 2023-10-14 22:57
 * FileName: LCP/LCR 025. 两数相加 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	type ListNode = lib.ListNode
	stack1, stack2 := []int{}, []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	sum, carry := []int{}, 0
	for pos1, pos2 := len(stack1)-1, len(stack2)-1; pos1 >= 0 || pos2 >= 0 || carry > 0; {
		if pos1 >= 0 {
			carry += stack1[pos1]
			pos1--
		}
		if pos2 >= 0 {
			carry += stack2[pos2]
			pos2--
		}
		sum = append([]int{carry % 10}, sum...)
		carry /= 10
	}
	dummy := &ListNode{Val: -1}
	cur := dummy
	for _, num := range sum {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	l1, l2 := lib.Str2List("[7,2,4,3]"), lib.Str2List("[5,6,4]")
	fmt.Println(lib.List2String(addTwoNumbers(l1, l2)))
}
