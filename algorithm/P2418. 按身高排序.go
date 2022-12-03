/**
 * Author: Deean
 * Date: 2022/12/3 21:44
 * FileName: algorithm/P2418. 按身高排序.go
 * Description:
 */

package main

import "fmt"

func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			if heights[i] < heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
				names[i], names[j] = names[j], names[i]
			}
		}
	}
	return names
}

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}
