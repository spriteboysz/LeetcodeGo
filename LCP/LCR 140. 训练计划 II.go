/**
 * Author: Deean
 * Date: 2023-10-16 22:59
 * FileName: LCP/LCR 140. 训练计划 II.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func trainingPlan2(head *lib.ListNode, cnt int) *lib.ListNode {
	slow, fast := head, head
	for cnt > 0 {
		fast = fast.Next
		cnt--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	fmt.Println(lib.List2String(trainingPlan2(lib.Str2List("[2,4,7,8]"), 1)))
}
