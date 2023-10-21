/**
 * Author: Deean
 * Date: 2023-10-19 23:41
 * FileName: LCR/LCR 182. 动态口令.go
 * Description:
 */

package main

import "fmt"

func dynamicPassword(password string, target int) string {
	return password[target:] + password[:target]
}

func main() {
	fmt.Println(dynamicPassword("s3cur1tyC0d3", 4))
}
