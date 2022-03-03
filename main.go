package main

import "fmt"

func say() string {
	return "Hello, Gopher"
}

func say2(s string) (string, int) {
	return s + " Hello, Gopher", 2
}

func main() {
	s := "Hello, Gopher"
	var age int
	fmt.Println(s)
	fmt.Println(age)

	h := say()
	fmt.Println(h)

	ha, hb := say2("foo")
	fmt.Println(ha, hb)
}