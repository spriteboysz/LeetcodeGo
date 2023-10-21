/**
 * Author: Deean
 * Date: 2023-10-19 23:43
 * FileName: LCR/LCR 189. 设计机械累加器.go
 * Description:
 */

package main

import "fmt"

func mechanicalAccumulator(target int) int {
	if target == 0 {
		return 0
	}
	return mechanicalAccumulator(target-1) + target
}

func main() {
	fmt.Println(mechanicalAccumulator(5))
	fmt.Println(mechanicalAccumulator(7))
}
