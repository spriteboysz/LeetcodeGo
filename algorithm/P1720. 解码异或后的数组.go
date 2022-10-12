/**
 * Author: Deean
 * Date: 2022-10-12 23:40
 * FileName: algorithm/P1720. 解码异或后的数组.go
 * Description:
 */

package main

import "fmt"

func decode(encoded []int, first int) []int {
	decoded := []int{first}
	for i, code := range encoded {
		decoded = append(decoded, decoded[i]^code)
	}
	return decoded
}

func main() {
	encoded := []int{6, 2, 7, 3}
	fmt.Println(decode(encoded, 4))
}
