/**
 * Author: Deean
 * Date: 2022/11/23 22:54
 * FileName: algorithm/P1736. 替换隐藏数字得到的最晚时间.go
 * Description:
 */

package main

import "fmt"

func maximumTime(time string) string {
	for i := 24*60 - 1; i >= 0; i-- {
		hh, mm := i/60, i%60
		cur := fmt.Sprintf("%02d:%02d", hh, mm)
		flag := true
		for j := 0; j < 5; j++ {
			if time[j] != '?' && cur[j] != time[j] {
				flag = false
				break
			}
		}
		if flag {
			return cur
		}
	}
	return ""
}

func main() {
	fmt.Println(maximumTime("0?:3?"))
}
