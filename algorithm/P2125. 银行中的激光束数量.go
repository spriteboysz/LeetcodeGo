/**
 * Author: Deean
 * Date: 2022-10-14 22:30
 * FileName: algorithm/P2125. 银行中的激光束数量.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func numberOfBeams(bank []string) int {
	var rows []int
	for _, s := range bank {
		cur := strings.Count(s, "1")
		if cur > 0 {
			rows = append(rows, cur)
		}
	}
	beams := 0
	for i := 1; i < len(rows); i++ {
		beams += rows[i-1] * rows[i]
	}
	return beams
}

func main() {
	bank := []string{"011001", "000000", "010100", "001000"}
	fmt.Println(numberOfBeams(bank))
}
