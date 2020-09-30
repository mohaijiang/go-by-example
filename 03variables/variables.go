package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	// 不定义 参数的初始化值，默认输出是什么，int ？
	var e string
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
