package main

import (
	"fmt"
	"time"
)

func hello_go() {
	fmt.Println("hello go!!!")
}

func main() {
	go hello_go()
	fmt.Println("main done!!!")
	time.Sleep(time.Second)
}

// 定义接口
type Sayer interface {
	say()
}

// 定义结构体
type dog struct {
}
type cat struct {
}

// 定义方法
func (d dog) say() {
	fmt.Println("狗叫")
}
func (c cat) say() {
	fmt.Println("猫叫")
}
