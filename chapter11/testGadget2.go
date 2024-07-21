package main

import (
	"fmt"
	mypkg "github.com/musienko-maxim/GoFirstBook/mypkg"
)

func main() {
	var value mypkg.MyType
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())

}
