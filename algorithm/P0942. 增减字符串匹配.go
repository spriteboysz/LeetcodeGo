/**
 * Author: Deean
 * Date: 2022-10-16 23:51
 * FileName: algorithm/P0942. 增减字符串匹配.go
 * Description:
 */

package main

import "fmt"

func diStringMatch(s string) []int {
	left, right := 0, len(s)
	var ret []int
	for _, c := range s {
		if c == 'I' {
			ret = append(ret, left)
			left++
		} else {
			ret = append(ret, right)
			right--
		}
	}
	ret = append(ret, left)
	return ret
}

func main() {
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
}
