/**
 * Author: Deean
 * Date: 2022/11/22 22:50
 * FileName: algorithm/P0326. 3 的幂.go
 * Description:
 */

package main

import "fmt"

func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(45))
}
