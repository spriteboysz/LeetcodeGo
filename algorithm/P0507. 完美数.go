/**
 * Author: Deean
 * Date: 2022/11/26 21:56
 * FileName: algorithm/P0507. 完美数.go
 * Description:
 */

package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	factors := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			factors += i
			if i*i < num {
				factors += num / i
			}
		}
	}
	return factors == num
}

func main() {
	fmt.Println(checkPerfectNumber(28))
}
