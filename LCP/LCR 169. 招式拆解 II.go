/**
 * Author: Deean
 * Date: 2023-10-19 22:45
 * FileName: LCR/LCR 169. 招式拆解 II.go
 * Description:
 */

package main

import "fmt"

func dismantlingAction(arr string) byte {
	hash := map[rune]int{}
	for _, c := range arr {
		hash[c]++
	}
	for _, c := range arr {
		if hash[c] == 1 {
			return byte(c)
		}
	}
	return ' '
}

func main() {
	fmt.Println(dismantlingAction("abbccdeff"))
}
