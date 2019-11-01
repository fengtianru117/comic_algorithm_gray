package main

import (
	"encoding/json"
	"fmt"
)

// https://www.jianshu.com/p/78b8c95dedc8
// empty Slice 和 nil Slice 的区别

func main() {
	var a []int
	fmt.Printf("%p\n", a)
	fmt.Println(a == nil)
	fmt.Println(a)
	fmt.Printf("%d\n", a)
	fmt.Println("#################")
	var b []int
	b = []int{}
	fmt.Printf("%p\n", b)
	fmt.Println(b == nil)
	fmt.Println(b)
	fmt.Printf("%d\n", b)
	fmt.Println("#########################################################")
	fmt.Println("#########################################################")
	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	fmt.Println(string(aj), string(bj))
}
