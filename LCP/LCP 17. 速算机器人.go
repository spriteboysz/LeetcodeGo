/**
 * Author: Deean
 * Date: 2022-10-16 16:58
 * FileName: LCP/LCP 17. 速算机器人.go
 * Description:
 */

package main

import "fmt"

func calculate(s string) int {
	x, y := 1, 0
	for _, c := range s {
		if c == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}

func main() {
	fmt.Println(calculate("AB"))
}
