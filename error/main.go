package main

import (
	"./error"
	"fmt"
)

func main() {
	err := error.Sample()
	if err != nil {
		fmt.Println(err)
	}
}
