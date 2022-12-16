/**
 * Author: Deean
 * Date: 2022/12/16 22:37
 * FileName: algorithm/P1227. 飞机座位分配概率.go
 * Description:
 */

package main

import "fmt"

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	}
	return 0.5
}

func main() {
	fmt.Println(nthPersonGetsNthSeat(1))
	fmt.Println(nthPersonGetsNthSeat(2))
}
