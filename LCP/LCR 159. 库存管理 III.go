/**
 * Author: Deean
 * Date: 2023-10-18 00:03
 * FileName: LCR/LCR 159. 库存管理 III.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func inventoryManagement(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}

func main() {
	fmt.Println(inventoryManagement([]int{0, 2, 3, 6}, 2))
}
