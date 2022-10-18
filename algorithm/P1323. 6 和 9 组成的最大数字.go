/**
 * Author: Deean
 * Date: 2022-10-18 23:09
 * FileName: algorithm/P1323. 6 和 9 组成的最大数字.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	num, _ = strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return num
}

func main() {
	fmt.Println(maximum69Number(9669))
}
