/**
 * Author: Deean
 * Date: 2023-10-19 23:37
 * FileName: LCR/LCR 179. 查找总价格为目标值的两个商品.go
 * Description:
 */

package main

import "fmt"

func twoSum2(price []int, target int) []int {
	left, right := 0, len(price)-1
	for left < right {
		if price[left]+price[right] < target {
			left++
		} else if price[left]+price[right] > target {
			right--
		} else {
			return []int{price[left], price[right]}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum2([]int{3, 9, 12, 15}, 18))
}
