package main

import "fmt"
import "strconv"

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


   // 配列
   array := [...]int{1, 3, 5}
   fmt.Println(array)
   // [1 3 5]
   fmt.Println(len(array))
   // 3
}

func sampleFunc(name string) (/* ここに帰り値を記述する事によりreturnの後にreturnする変数名を記述しなくていい */ msg string) {
    msg = "Input name is " + name
    return
}
