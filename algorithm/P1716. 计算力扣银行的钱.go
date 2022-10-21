/**
 * Author: Deean
 * Date: 2022-10-21 21:19
 * FileName: algorithm/P1716. 计算力扣银行的钱.go
 * Description:
 */

package main

import "fmt"

func totalMoney(n int) int {
	total, week, day := 0, 1, 1
	for i := 0; i < n; i++ {
		total += week + day - 1
		day++
		if day == 8 {
			day = 1
			week++
		}
	}
	return total
}

func main() {
	fmt.Println(totalMoney(20))
}
