/**
 * Author: Deean
 * Date: 2023-10-16 22:46
 * FileName: LCP/LCR 128. 库存管理 I.go
 * Description:
 */

package main

import "fmt"

func stockManagement(stock []int) int {
	minimum := stock[0]
	for _, num := range stock {
		if minimum > num {
			minimum = num
		}
	}
	return minimum
}

func main() {
	fmt.Println(stockManagement([]int{5, 7, 9, 1, 2}))
}
