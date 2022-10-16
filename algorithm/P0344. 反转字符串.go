/**
 * Author: Deean
 * Date: 2022-10-16 19:59
 * FileName: algorithm/P0344. 反转字符串.go
 * Description:
 */

package main

import "fmt"

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	fmt.Println(s)
}

func main() {
	s := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(s)
}
