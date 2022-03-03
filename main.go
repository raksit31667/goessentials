package main

import (
	"fmt"

	"net/http"

	"github.com/labstack/echo"
)

type Person struct {
	Name string
	Age int
}

func (p Person) Walk() {
	fmt.Println(p.Name, "is walking")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		var p Person = Person{
			Age: 30,
			Name: "John",
		}

		fmt.Printf("%#v\n", p)

		p.Walk()

		return c.String(http.StatusOK, "Hello, Gopher")
	})

	e.Logger.Fatal(e.Start(":1234"))
}