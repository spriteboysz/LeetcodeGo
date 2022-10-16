/**
 * Author: Deean
 * Date: 2022-10-16 18:01
 * FileName: algorithm/P1837. K 进制表示下的各位数字总和.go
 * Description:
 */

package main

import "fmt"

func sumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}

func main() {
	fmt.Println(sumBase(34, 6))
}
