/**
 * Author: Deean
 * Date: 2022/11/15 22:18
 * FileName: algorithm/P2469. 温度转换.go
 * Description:
 */

package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.0}
}

func main() {
	fmt.Println(convertTemperature(122.11))
}
