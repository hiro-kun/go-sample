package main

import (
    "fmt"
    "time"
    "strconv"
)

// 構造体
type product struct {
    id string
    price int
}

// 値渡し
func (p product) show() {
    fmt.Println(p)
}
// 参照渡し
func (p *product) add(value int) {
    p.price+= value
}

// インターフェイス
type animal interface {
    eat()
}

// 空インターフェイス
func showAnimalType(t interface{}) {
    /*
    _, isDog := t.(dog)

    if isDog {
        fmt.Println("Dog.")
    } else {
        fmt.Println("Not Dog.")
    }
    */

    // 型Switch
    switch t.(type) {
    case dog:
        fmt.Println("Dog.")
    case cat:
        fmt.Println("Cat.")
    default:
        fmt.Println("Not animal.")
    }
}

type cat struct {}
type dog struct {}

func (c cat) eat() {
    fmt.Println("Cat eat!")
}
func (d dog) eat() {
    fmt.Println("Dog eat!")
}

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
    /*
        0 a
        1 b
        2 c
    */
    // ブランク修飾子（_）を利用する事によりvalueのみ引き出しが可能
    for _, value := range rangeSample {
        fmt.Println(value)
    }
    /*
        a
        b
        c
    */
    // mapでのrange
    rangeMapSample := map[string]int{"yamada":100, "satou":200, "suzuki":300}
    for key, value := range rangeMapSample {
        fmt.Println(key, value)
    }
    /*
        suzuki 300
        yamada 100
        satou 200
    */

    // 構造体利用
    product := new(product)
    product.id = "A0001"
    product.price = 200
    fmt.Println(product)
    // &{A0001 200}

    product.show()
    // {A0001 200}
    product.add(500)
    product.show()
    // {A0001 700}


    // インターフェイス ※他の型の場合、rangeでループ出来なくなるが、インターフェイスをあてる事によりループ処理を回す事が出来る
    animals := []animal{cat{}, dog{}}
       for _, animal := range animals {
       animal.eat()
       showAnimalType(animal)
       // Cat eat!
       // Cat.
       // Dog eat!
       // Dog.
    }

    // goルーチン
    Isfinished := make(chan bool)

    // 配列
    execs := []func(){
        func() {
            fmt.Println("exec1 started.")
            time.Sleep(1 * time.Second)
            fmt.Println("exec1 finished.")
            Isfinished <- true
        },
        func() {
            fmt.Println("exec2 started.")
            time.Sleep(2 * time.Second)
            fmt.Println("exec2 finished.")
            Isfinished <- true
        },
        func() {
            fmt.Println("exec3 started.")
            time.Sleep(3 * time.Second)
            fmt.Println("exec3 finished.")
            Isfinished <- true
        },
    }

    // 並行化
    for _, exec := range execs {
        go exec()
    }

    // 終了待ち
    for i := 0; i < len(execs); i++ {
        <-Isfinished
    }

    fmt.Println("all finished.")
}

func sampleFunc(name string) (/* ここに帰り値を記述する事によりreturnの後にreturnする変数名を記述しなくていい */ msg string) {
    msg = "Input name is " + name
    return
}
