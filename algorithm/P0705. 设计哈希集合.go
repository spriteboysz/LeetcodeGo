/**
 * Author: Deean
 * Date: 2022/11/13 16:25
 * FileName: algorithm/P0705. 设计哈希集合.go
 * Description:
 */

package main

import (
	"fmt"
)

type TypeName map[int]interface{}

type MyHashSet struct {
	TypeName
}

func Constructor2() MyHashSet {
	var myHash = make(map[int]interface{})

	return MyHashSet{
		myHash,
	}
}

func (this *MyHashSet) Add(key int) {
	this.TypeName[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	delete(this.TypeName, key)
}

func (this *MyHashSet) Contains(key int) bool {
	if this.TypeName[key] != nil {
		return true
	} else {
		return false
	}
}

func main() {
	myHashSet := Constructor2()
	myHashSet.Add(1)                   // set = [1]
	myHashSet.Add(2)                   // set = [1, 2]
	fmt.Println(myHashSet.Contains(1)) // 返回 True
	fmt.Println(myHashSet.Contains(3)) // 返回 False ，（未找到）
	myHashSet.Add(2)                   // set = [1, 2]
	fmt.Println(myHashSet.Contains(2)) // 返回 True
	myHashSet.Remove(2)                // set = [1]
	fmt.Println(myHashSet.Contains(2)) // 返回 False ，（已移除）
}
