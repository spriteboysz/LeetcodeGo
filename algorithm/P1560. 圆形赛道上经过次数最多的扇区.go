/**
 * Author: Deean
 * Date: 2022/11/20 17:40
 * FileName: algorithm/P1560. 圆形赛道上经过次数最多的扇区.go
 * Description:
 */

package main

import "fmt"

func mostVisited(n int, rounds []int) []int {
	a, b := rounds[0], rounds[len(rounds)-1]
	r := []int{}
	if a <= b {
		for i := a; i <= b; i++ {
			r = append(r, i)
		}
	} else {
		for i := 1; i <= b; i++ {
			r = append(r, i)
		}
		for i := a; i <= n; i++ {
			r = append(r, i)
		}
	}
	return r
}

func main() {
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
}
