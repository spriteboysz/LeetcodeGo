/**
 * Author: Deean
 * Date: 2022-10-16 21:13
 * FileName: algorithm/P1528. 重新排列字符串.go
 * Description:
 */

package main

import "fmt"

func restoreString(s string, indices []int) string {
	ss := make([]byte, len(s))
	for i, index := range indices {
		ss[index] = s[i]
	}
	return string(ss)
}

func main() {
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString("codeleet", indices))
}
