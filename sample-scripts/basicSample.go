package main

import "fmt"

func main() {
    // 変数代入
    msg := "hello world"
    fmt.Println(msg)
    // hello world

    num := 12345
    fmt.Println(num)
    // 12345


    // 定数
    const sampleConst = "sample message"
    fmt.Println(sampleConst)
    // sample message


    // 関数
    fmt.Println(sampleFunc("yamada-tarou"))
    // Input name is yamada-tarou
}

func sampleFunc(name string) (/* ここに帰り値を記述する事によりreturnの後にreturnする変数名を記述しなくていい */ msg string) {
    msg = "Input name is " + name
    return
}
