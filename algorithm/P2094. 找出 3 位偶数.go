/**
 * Author: Deean
 * Date: 2022/11/25 22:05
 * FileName: algorithm/P2094. 找出 3 位偶数.go
 * Description:
 */

package main

import "fmt"

func findEvenNumbers(digits []int) []int {
	numbers := [10]int{}
	for _, digit := range digits {
		numbers[digit]++
	}
	check := func(num int) bool {
		cur := [10]int{}
		for i := 0; i <= 9; i++ {
			cur[i] = numbers[i]
		}
		a, b, c := num/100, num%100/10, num%10
		cur[a]--
		cur[b]--
		cur[c]--
		for i := 0; i <= 9; i++ {
			if cur[i] < 0 {
				return false
			}
		}
		return true
	}
	even := []int{}
	for i := 100; i < 999; i += 2 {
		if check(i) {
			even = append(even, i)
		}
	}
	return even
}

func main() {
	fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))
}
