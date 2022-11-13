/**
 * Author: Deean
 * Date: 2022/11/13 16:48
 * FileName: algorithm/P1491. 去掉最低工资和最高工资后的工资平均值.go
 * Description:
 */

package main

import "fmt"

func average(salary []int) float64 {
	total, minimum, maximum := 0, 1000000, 0
	for _, num := range salary {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
		total += num
	}
	return float64(total-maximum-minimum) / float64(len(salary)-2)
}

func main() {
	fmt.Println(average([]int{6000, 5000, 4000, 3000, 2000, 1000}))
}
