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


   // スライス
   hoge := []int{1, 2, 3}
   fmt.Println(hoge)
   // [1 2 3]
   hoge = append(hoge, 4, 5)
   fmt.Println(hoge)
   // [1 2 3 4 5]


   // マップ
   mapSampleA := map[string]int{"yamada":30, "tarou":35}
   fmt.Println(mapSampleA)
   // map[yamada:30 tarou:35]

   mapStringSampleB := map[string]string{"ishikawa":"kanazawa"}
   fmt.Println(mapStringSampleB)
   // map[ishikawa:kanazawa]

   delete(mapSampleA, "yamada")
   fmt.Println(mapSampleA)
   // map[tarou:35]

   value, isExist := mapSampleA["tarou"]
   fmt.Println(value)
   // 35
   fmt.Println(isExist)
   // true


   // if文
   point := 1
   if point % 3 == 0 && point % 5 == 0 {
       fmt.Println("FizzBuzz")
   } else if point % 5 == 0 {
       fmt.Println("Fizz")
   } else if point % 5 == 0 {
       fmt.Println("Buzz")
   } else {
       fmt.Println("else")
   }


   // switch文
   country := "France"
       switch country {
       case "Japan":
           fmt.Println("Japan")
       case "UK", "France", "Italy":
           fmt.Println("Europe")
       default:
           fmt.Println("Not applicable")
    }

    height := 150
    switch {
    case height > 175:
            fmt.Println("tall")
        default:
            fmt.Println("common people")
    }


    // for文
    for i := 0; i < 5; i++ {
        // ループから抜ける
        if i == 2 { break }
        // 処理スキップ
        if i == 4 { continue }
        fmt.Println(i)
    }

    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    j := 0
    for {
        fmt.Println(j)
        j++
        if j == 5 { break }
    }


    // range
    rangeSample := []string{"a", "b", "c"}
    for key, value := range rangeSample {
        fmt.Println(key, value)
    }
    // ブランク修飾子（_）を利用する事によりvalueのみ引き出しが可能
    for _, value := range rangeSample {
        fmt.Println(value)
    }
    // mapでのrange
    rangeMapSample := map[string]int{"yamada":100, "satou":200, "suzuki":300}
    for key, value := range rangeMapSample {
        fmt.Println(key, value)
    }
}

func sampleFunc(name string) (/* ここに帰り値を記述する事によりreturnの後にreturnする変数名を記述しなくていい */ msg string) {
    msg = "Input name is " + name
    return
}
