/**
 * Author: Deean
 * Date: 2022-10-16 21:05
 * FileName: algorithm/P0657. 机器人能否返回原点.go
 * Description:
 */

package main

import "fmt"

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			y += 1
		case 'D':
			y -= 1
		case 'L':
			x -= 1
		case 'R':
			x += 1
		}
	}
	return x == 0 && y == 0
}

func main() {
	fmt.Println(judgeCircle("LLRURUDD"))
}
