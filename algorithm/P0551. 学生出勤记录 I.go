/**
 * Author: Deean
 * Date: 2022/11/26 22:11
 * FileName: algorithm/P0551. 学生出勤记录 I.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
