/**
 * Author: Deean
 * Date: 2023-10-19 22:57
 * FileName: LCR/LCR 173. 点名.go
 * Description:
 */

package main

import "fmt"

func takeAttendance(records []int) int {
	for i, record := range records {
		if i != record {
			return i
		}
	}
	return len(records)
}

func main() {
	fmt.Println(takeAttendance([]int{0, 1, 2, 3, 4, 5, 6, 8}))
	fmt.Println(takeAttendance([]int{0, 1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(takeAttendance([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
