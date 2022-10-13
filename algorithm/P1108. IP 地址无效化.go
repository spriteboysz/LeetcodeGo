/**
 * Author: Deean
 * Date: 2022-10-13 22:55
 * FileName: algorithm/P1108. IP 地址无效化.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}
