package main

import (
	"fmt"
	"github.com/raksit31667/goessentials/say"
)

func main() {
	s := "Hello, Gopher"
	var age int
	fmt.Println(s)
	fmt.Println(age)

	h := say.Say()
	fmt.Println(h)

	ha, hb := say.Say2("foo")
	fmt.Println(ha, hb)
}