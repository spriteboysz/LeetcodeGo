/**
 * Author: Deean
 * Date: 2022/12/1 22:52
 * FileName: algorithm/P2481. 分割圆的最少切割次数.go
 * Description:
 */

package main

import "fmt"

func numberOfCuts(n int) int {
	if n == 1 || n%2 == 0 {
		return n / 2
	}
	return n
}

func main() {
	fmt.Println(numberOfCuts(4))
}
