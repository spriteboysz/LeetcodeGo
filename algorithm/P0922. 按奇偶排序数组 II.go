/**
 * Author: Deean
 * Date: 2022/11/12 14:52
 * FileName: algorithm/P0922. 按奇偶排序数组 II.go
 * Description:
 */

package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	even, odd, arr := []int{}, []int{}, []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	for i := 0; i < len(even); i++ {
		arr = append(arr, even[i])
		arr = append(arr, odd[i])
	}
	return arr
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
