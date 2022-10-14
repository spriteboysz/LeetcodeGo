/**
 * Author: Deean
 * Date: 2022-10-14 21:27
 * FileName: LCP/LCP 01. 猜数字.go
 * Description:
 */

package main

import "fmt"

func game(guess []int, answer []int) int {
	cnt := 0
	for i, g := range guess {
		if g == answer[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	guess, answer := []int{2, 2, 3}, []int{3, 2, 1}
	fmt.Println(game(guess, answer))
}
