/**
 * Author: Deean
 * Date: 2022/11/12 20:46
 * FileName: algorithm/P1518. 换水问题.go
 * Description:
 */

package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	bottles, total := numBottles, numBottles
	for bottles >= numExchange {
		bottles -= numExchange
		total += 1
		bottles += 1
	}
	return total
}

func main() {
	fmt.Println(numWaterBottles(15, 4))
}
