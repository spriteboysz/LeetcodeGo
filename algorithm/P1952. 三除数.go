/**
 * Author: Deean
 * Date: 2022/11/25 22:19
 * FileName: algorithm/P1952. 三除数.go
 * Description:
 */

package main

import "fmt"

func isThree(n int) bool {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return cnt == 3
}

func main() {
	fmt.Println(isThree(4))
}
