/**
 * Author: Deean
 * Date: 2022-10-14 21:54
 * FileName: algorithm/P1688. 比赛中的配对次数.go
 * Description:
 */

package main

import "fmt"

func numberOfMatches(n int) int {
	cnt := 0
	for n > 1 {
		cnt += n / 2
		if n%2 == 0 {
			n /= 2
		} else {
			n = n/2 + 1
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfMatches(14))
}
