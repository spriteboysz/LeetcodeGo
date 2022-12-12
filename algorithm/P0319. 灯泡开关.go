/**
 * Author: Deean
 * Date: 2022/12/11 22:33
 * FileName: algorithm/P0319. 灯泡开关.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}

func main() {
	fmt.Println(bulbSwitch(0))
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(2))
	fmt.Println(bulbSwitch(3))
}
