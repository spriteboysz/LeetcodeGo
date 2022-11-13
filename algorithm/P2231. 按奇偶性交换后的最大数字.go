/**
 * Author: Deean
 * Date: 2022/11/13 17:29
 * FileName: algorithm/P2231. 按奇偶性交换后的最大数字.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func largestInteger(num int) int {
	even, odd, numbers := []int{}, []int{}, []int{}
	for ; num > 0; num /= 10 {
		if num%10%2 == 0 {
			even = append(even, num%10)
		} else {
			odd = append(odd, num%10)
		}
		numbers = append(numbers, num%10%2)
	}
	sort.Slice(even, func(i, j int) bool {
		return even[i] > even[j]
	})
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})
	j, k, largest := 0, 0, 0
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] == 0 {
			largest = largest*10 + even[j]
			j++
		} else {
			largest = largest*10 + odd[k]
			k++
		}
	}
	return largest
}

func main() {
	fmt.Println(largestInteger(65875))
}
