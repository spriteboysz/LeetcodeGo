/**
 * Author: Deean
 * Date: 2022-10-21 21:48
 * FileName: algorithm/P1189. “气球” 的最大数量.go
 * Description:
 */

package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	count := [5]float32{}
	for _, c := range text {
		switch c {
		case 'b':
			count[0]++
		case 'a':
			count[1]++
		case 'l':
			count[2] += 0.5
		case 'o':
			count[3] += 0.5
		case 'n':
			count[4]++
		}
	}
	cnt := len(text)
	for _, c := range count {
		if cnt > int(c) {
			cnt = int(c)
		}
	}
	return cnt
}

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}
