/**
 * Author: Deean
 * Date: 2022/12/13 23:12
 * FileName: algorithm/P1417. 重新格式化字符串.go
 * Description:
 */

package main

import "fmt"

func reformat(s string) string {
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	letter, digit, ss := []byte{}, []byte{}, []byte{}
	for _, c := range s {
		if c < 'A' {
			digit = append(digit, byte(c))
		} else {
			letter = append(letter, byte(c))
		}
	}
	if abs(len(letter)-len(digit)) > 1 {
		return ""
	}
	if len(letter) < len(digit) {
		letter, digit = digit, letter
	}
	for i := 0; i < len(digit); i++ {
		ss = append(append(ss, letter[i]), digit[i])
	}
	if len(digit) != len(letter) {
		ss = append(ss, letter[len(letter)-1])
	}
	return string(ss)
}

func main() {
	fmt.Println(reformat("covid2019"))
}
