package main

import (
	"fmt"
	"strconv"
)

// 基本的な関数
func sampleFunc(name string) ( /* ここに帰り値を記述する事によりreturnの後にreturnする変数名を記述しなくていい */ msg string) {
	msg = "Input name is " + name
	return
}

func main() {

	// 関数代入
	f := func(name string, age int) (msg string) {
		msg = "His name is " + name + " and his age is " + strconv.Itoa(age)
		return
	}
	fmt.Println(f("yamada-tarou", 30))
	// His name is yamada-tarou and his age is 30

	// 即時関数
	func(msg string) {
		fmt.Println(msg)
	}("func test")
	// func test

	fmt.Println(sampleFunc("yamada-tarou"))
	// Input name is yamada-tarou
}
