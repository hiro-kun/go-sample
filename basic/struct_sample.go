package main

import (
	"fmt"
)

// 構造体
type product struct {
	id    string
	price int
}

// 値渡し
func (p product) show() {
	fmt.Println(p)
}

// 参照渡し
func (p *product) add(value int) {
	p.price += value
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

type cat struct{}
type dog struct{}

func (c cat) eat() {
	fmt.Println("Cat eat!")
}
func (d dog) eat() {
	fmt.Println("Dog eat!")
}

func main() {
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
}
