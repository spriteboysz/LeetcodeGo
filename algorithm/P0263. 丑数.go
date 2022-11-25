/**
 * Author: Deean
 * Date: 2022/11/25 22:00
 * FileName: algorithm/P0263. 丑数.go
 * Description:
 */

package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(1))
}
