- While learn Go, focus on engineering practice rather than syntax
- function main in package main
- public method starts with uppercase, private mthod starts with lowercase
- Go has zero value
- No matter how many files in the package, Go considers each file as a package
- Test file name convention: xxx_test.go
- Go doesn't allow to declare var without using it, use _ instead
- go mod tidy to install dependencies, it generate go.sum (yarn.lock)
- Go doesn't accept slice (array) with different specific length on function argument
- Go requires comma on struct declaration so that we won't have to remove comma on last element

```go
var p Person = Person{
			Age: 30,
			Name: "John",
		}
```