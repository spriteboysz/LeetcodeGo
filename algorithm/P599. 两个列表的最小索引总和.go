/**
 * Author: Deean
 * Date: 2022/11/20 17:49
 * FileName: algorithm/P599. 两个列表的最小索引总和.go
 * Description:
 */

package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	list := map[string]int{}
	index1, index2 := map[string]int{}, map[string]int{}
	for i, r := range list1 {
		list[r]++
		index1[r] = i
	}
	for i, r := range list2 {
		list[r]++
		index2[r] = i
	}
	minimum, target := 4002, []string{}
	for k, v := range list {
		if v == 2 && index1[k]+index2[k] < minimum {
			minimum = index1[k] + index2[k]
		}
	}

	for k, v := range list {
		if v == 2 && index1[k]+index2[k] == minimum {
			target = append(target, k)
		}
	}
	return target
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"}))
}
