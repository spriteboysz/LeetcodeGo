/**
 * Author: Deean
 * Date: 2022/12/10 22:59
 * FileName: algorithm/P1807. 替换字符串中的括号内容.go
 * Description:
 */

package main

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	hash := map[string]string{}
	for _, pair := range knowledge {
		hash[pair[0]] = pair[1]
	}
	ss, n := []byte{}, len(s)
	for i := 0; i < n; i++ {
		key := ""
		if s[i] == '(' {
			j := i + 1
			for ; j < n && s[j] != ')'; j++ {
			}
			key = s[i+1 : j]
			val, ok := hash[key]
			if ok {
				for k := 0; k < len(val); k++ {
					ss = append(ss, val[k])
				}
			} else {
				ss = append(ss, '?')
			}
			i = j
		} else {
			ss = append(ss, s[i])
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}))
}
