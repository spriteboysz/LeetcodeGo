/**
 * Author: Deean
 * Date: 2022/11/15 23:32
 * FileName: algorithm/P0806. 写字符串需要的行数.go
 * Description:
 */

package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	lines, width, maxWidth := 1, 0, 100
	for _, c := range s {
		need := widths[c-'a']
		width += need
		if width > maxWidth {
			lines++
			width = need
		}
	}
	return []int{lines, width}
}

func main() {
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		"bbbcccdddaaa"))
}
