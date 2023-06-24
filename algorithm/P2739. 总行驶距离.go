/**
 * Author: Deean
 * Date: 2023-06-24 15:32
 * FileName: algorithm/P2739. 总行驶距离.go
 * Description:
 */

package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	return (mainTank + min((mainTank-1)/4, additionalTank)) * 10
}

func main() {
	fmt.Println(distanceTraveled(5, 10))
}
