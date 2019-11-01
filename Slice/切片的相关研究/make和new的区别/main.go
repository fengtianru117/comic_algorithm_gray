package main

import (
	"fmt"
	"sync"
)

// https://www.flysnow.org/2017/10/23/go-new-vs-make.html

type user struct{
	lock sync.Mutex
	name *string
	age int
}

func main(){
	a:=new(user)
	fmt.Printf("%+v\n",a)
	b:=make([]user,1,1)
	fmt.Printf("%+v\n",b)
	// 引用了nil的地址，会报错
	fmt.Println(*b[0].name)
	c:=make([]user,0)
	fmt.Printf("%+v\n",c)
}
