/**
 * Author: Deean
 * Date: 2022-10-17 23:23
 * FileName: algorithm/P1309. 解码字母到整数映射.go
 * Description:
 */

package main

import "fmt"

func freqAlphabets(s string) string {
	var ss []byte
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			ss = append(ss, (s[i]-'0')*10+s[i+1]-'0'+'a'-1)
			i += 2
		} else {
			ss = append(ss, s[i]-'0'+'a'-1)
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(freqAlphabets("1326#10#11#12"))
}
