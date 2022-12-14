/**
 * Author: Deean
 * Date: 2022/12/14 22:13
 * FileName: LCP/LCS 01. 下载插件.go
 * Description:
 */

package main

import "fmt"

func leastMinutes(n int) int {
	if n < 2 {
		return n
	}

	for i, s := 1, 1; i < n; i++ {
		s *= 2
		if s >= n {
			return i + 1
		}
	}
	return n
}

func main() {
	fmt.Println(leastMinutes(2))
	fmt.Println(leastMinutes(4))
}
