/**
 * Author: Deean
 * Date: 2023-10-14 20:44
 * FileName: LCP/LCR 003. 比特位计数.go
 * Description:
 */

package main

import "fmt"

func countBits(n int) []int {
	bits := []int{}
	for i := 0; i <= n; i += 1 {
		num, cnt := i, 0
		for num > 0 {
			num &= num - 1
			cnt += 1
		}
		bits = append(bits, cnt)
	}
	return bits
}

func main() {
	fmt.Println(countBits(5))
}
