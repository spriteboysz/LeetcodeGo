/**
 * Author: Deean
 * Date: 2022-10-15 23:49
 * FileName: algorithm/P1436. 旅行终点站.go
 * Description:
 */

package main

import "fmt"

func destCity(paths [][]string) string {
	var src []string
	var dst []string
	for _, path := range paths {
		src = append(src, path[0])
		dst = append(dst, path[1])
	}
	for _, d := range dst {
		flag := true
		for _, s := range src {
			if d == s {
				flag = false
				break
			}
		}
		if flag {
			return d
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"Ba", "Ca"}, {"Da", "Ba"}, {"Ca", "Aa"}}
	fmt.Println(destCity(paths))
}
