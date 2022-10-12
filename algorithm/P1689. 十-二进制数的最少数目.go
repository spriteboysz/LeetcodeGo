/**
 * Author: Deean
 * Date: 2022-10-12 23:27
 * FileName: algorithm/P1689. 十-二进制数的最少数目.go
 * Description:
 */

package main

import "fmt"

func minPartitions(n string) int {
	maximum := 0
	for _, num := range n {
		if int(num-'0') > maximum {
			maximum = int(num - '0')
		}
		if maximum == 9 {
			return maximum
		}
	}
	return maximum
}

func main() {
	fmt.Println(minPartitions("27346209830709182346"))
}
