package main

import "fmt"

// https://blog.csdn.net/youngwhz1/article/details/83026263

func sliceDelete(s []interface{}, index int) []interface{} {
	rear := append([]interface{}{}, s[index+1:]...)
	return append(s[:index], rear...)

}

func main() {
	var a = []interface{}{1, 2, 3, 4}
	// fmt.Println(a[:2])
	fmt.Println(sliceDelete(a, 2))
}
