package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) foo() {
	fmt.Println("foo")
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Println("i want to implement something different for employee")
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-2232"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "232-232-3921"}, "golang"}
	fmt.Printf("mark school %v", mark.school)
	mark.SayHi()
	sam.SayHi()
	mark.foo()
}
