package main

import "fmt"

//变量
func main() {
	//标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	//批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
}
