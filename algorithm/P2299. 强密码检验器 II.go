/**
 * Author: Deean
 * Date: 2022/11/19 16:05
 * FileName: algorithm/P2299. 强密码检验器 II.go
 * Description:
 */

package main

import (
	"fmt"
	"unicode"
)

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	flag1, flag2, flag3, flag4 := false, false, false, false
	for i, c := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		} else if unicode.IsDigit(c) {
			flag1 = true
		} else if unicode.IsUpper(c) {
			flag2 = true
		} else if unicode.IsLower(c) {
			flag3 = true
		} else {
			flag4 = true
		}
	}
	return flag1 && flag2 && flag3 && flag4
}

func main() {
	fmt.Println(strongPasswordCheckerII("IloveLe3tcode!"))
}
