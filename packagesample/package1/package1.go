package package1

import (
	"fmt"
  "strconv"
)

func UserInfo(name string, age int)(){
	msg := fmtUserInfo(name, age)
	fmt.Println(msg)
}

func fmtUserInfo(name string, age int)(string){
	msg := "Name is " + name + ", age is " + strconv.Itoa(age)

	return msg
}
