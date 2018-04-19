package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {

	// プログラム終了
	// os.Exit(0)

	// 変数代入
	msg := "hello world"
	fmt.Println(msg)
	// hello world

	num := 12345
	fmt.Println(num)
	// 12345

	// 定数
	const SampleConst = "sample message"
	fmt.Println(SampleConst)
	// sample message

	// 型取得
	fmt.Println(reflect.TypeOf(SampleConst))
	// string

	// 時刻比較
	nowTime := time.Now()
	// 一週間前の時刻に設定
	oneWeekAgo := nowTime.AddDate(0, 0, -7)
	// 時刻nowTimeは過去である。引数oneWeekAgoより。
	// ※同じ時刻は含まれない
	// false
	fmt.Println(nowTime.Before(oneWeekAgo))

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
	mapSampleA := map[string]int{"yamada": 30, "tarou": 35}
	fmt.Println(mapSampleA)
	// map[yamada:30 tarou:35]

	mapStringSampleB := map[string]string{"ishikawa": "kanazawa"}
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
	if point%3 == 0 && point%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if point%5 == 0 {
		fmt.Println("Fizz")
	} else if point%5 == 0 {
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
		if i == 2 {
			break
		}
		// 処理スキップ
		if i == 4 {
			continue
		}
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
		if j == 5 {
			break
		}
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
	rangeMapSample := map[string]int{"yamada": 100, "satou": 200, "suzuki": 300}
	for key, value := range rangeMapSample {
		fmt.Println(key, value)
	}
	/*
	   suzuki 300
	   yamada 100
	   satou 200
	*/

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

	// WEBサーバーを立てる
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
}
