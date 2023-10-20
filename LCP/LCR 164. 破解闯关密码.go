/**
 * Author: Deean
 * Date: 2023-10-18 22:30
 * FileName: LCR/LCR 164. 破解闯关密码.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func crackPassword(password []int) string {
	sort.Slice(password, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", password[i], password[j])
		y := fmt.Sprintf("%d%d", password[j], password[i])
		return x < y
	})

	res := ""
	for i := 0; i < len(password); i++ {
		res += fmt.Sprintf("%d", password[i])
	}
	return res
}

func main() {
	fmt.Println(crackPassword([]int{0, 3, 30, 34, 5, 9}))
}
