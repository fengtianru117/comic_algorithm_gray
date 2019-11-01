package main

import "fmt"

func main(){
	s:= []int{0:1,1:2,2:3,3:4,4:5}
	fmt.Println(cap(s))
	s[2]=10
	fmt.Println(s)
}
