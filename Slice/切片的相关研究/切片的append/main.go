package main

import "fmt"

func main(){
	a:=[]int{1,2,3,4,5}
	fmt.Println(cap(a))
	a=append(a,6)
	fmt.Println(cap(a))


	var b string
	fmt.Println(b=="")
	var c int
	fmt.Println(c)
}
