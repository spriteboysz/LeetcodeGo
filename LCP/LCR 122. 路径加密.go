/**
 * Author: Deean
 * Date: 2023-10-15 23:33
 * FileName: LCP/LCR 122. 路径加密.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func pathEncryption(path string) string {
	return strings.ReplaceAll(path, ".", " ")
}

func main() {
	fmt.Println(pathEncryption("a.aef.qerf.bb"))
}
