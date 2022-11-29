/**
 * Author: Deean
 * Date: 2022/11/29 23:33
 * FileName: algorithm/P0537. 复数乘法.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	process := func(num string) (real, imag int) {
		i := strings.IndexByte(num, '+')
		real, _ = strconv.Atoi(num[:i])
		imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
		return
	}
	real1, imag1 := process(num1)
	real2, imag2 := process(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+imag1*real2)
}

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}
