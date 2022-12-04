/**
 * Author: Deean
 * Date: 2022/12/4 17:02
 * FileName: algorithm/P1652. 拆炸弹.go
 * Description:
 */

package main

import "fmt"

func decrypt(code []int, k int) []int {
	n := len(code)
	decoded := make([]int, n)
	if k == 0 {
		return decoded
	}
	code = append(code, code...)
	left, right := 1, k
	if k < 0 {
		left, right = n+k, n-1
	}
	sum := 0
	for _, num := range code[left : right+1] {
		sum += num
	}
	for i := range decoded {
		decoded[i] = sum
		sum -= code[left]
		sum += code[right+1]
		left, right = left+1, right+1
	}
	return decoded
}

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
}
