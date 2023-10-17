/**
 * Author: Deean
 * Date: 2023-10-15 23:28
 * FileName: LCP/LCR 120. 寻找文件副本.go
 * Description:
 */

package main

import "fmt"

func findRepeatDocument(documents []int) int {
	hash := map[int]bool{}
	for _, document := range documents {
		if hash[document] {
			return document
		}
		hash[document] = true
	}
	return -1
}

func main() {
	fmt.Println(findRepeatDocument([]int{2, 5, 3, 0, 5, 0}))
}
