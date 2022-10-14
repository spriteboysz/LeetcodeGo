/**
 * Author: Deean
 * Date: 2022-10-14 21:22
 * FileName: algorithm/P1431. 拥有最多糖果的孩子.go
 * Description:
 */

package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maximum := 0
	for _, candy := range candies {
		if candy > maximum {
			maximum = candy
		}
	}
	var kids []bool
	for _, candy := range candies {
		if candy+extraCandies >= maximum {
			kids = append(kids, true)
		} else {
			kids = append(kids, false)
		}
	}
	return kids
}

func main() {
	candies := []int{12, 1, 12}
	fmt.Println(kidsWithCandies(candies, 10))
}
