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

type Animal struct {
	Name string
	Age int
}

type Walker interface {
	Walk()
}

type Runner interface {
	Run()
}

func (p Person) Walk() {
	fmt.Println(p.Name, "is walking")
}

func (p Person) Run() {
	fmt.Println(p.Name, "is running")
}

func (a Animal) Walk() {
	fmt.Println(a.Name, "is walking")
}

func (a Animal) Run() {
	fmt.Println(a.Name, "is running")
}

func PleaseWalk(w Walker) {
	w.Walk()
}

func PleaseRun(r Runner) {
	r.Run()
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

		var a Animal = Animal{Name: "Dog", Age: 22}

		a.Walk()

		PleaseWalk(p)
		PleaseWalk(a)
		PleaseRun(a)

		return c.String(http.StatusOK, "Hello, Gopher")
	})

	e.Logger.Fatal(e.Start(":1234"))
}