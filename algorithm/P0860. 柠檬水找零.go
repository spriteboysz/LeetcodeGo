/**
 * Author: Deean
 * Date: 2022/11/19 22:28
 * FileName: algorithm/P0860. 柠檬水找零.go
 * Description:
 */

package main

import "fmt"

func lemonadeChange(bills []int) bool {
	cnt05, cnt10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			cnt05++
		} else if bill == 10 {
			if cnt05 > 0 {
				cnt05--
				cnt10++
			} else {
				return false
			}
		} else {
			if cnt10 > 0 && cnt05 > 0 {
				cnt10--
				cnt05--
			} else if cnt05 > 2 {
				cnt05 -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}
