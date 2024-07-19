package main

import "fmt"

func main() {
	var x interface{}
	x = 123
	//v, ok := x.(int)
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
