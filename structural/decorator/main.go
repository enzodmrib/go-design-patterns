package main

import (
	"fmt"

	"github.com/enzodmrib/go-design-patterns/structural/decorator/decorator"
)

func main() {
	base := decorator.NewBase("Hi there!")
	impl1 := decorator.NewImpl1(base)
	fmt.Println(impl1.Render())
	impl2 := decorator.NewImpl2(impl1)
	fmt.Println(impl2.Render())
	impl3 := decorator.NewImpl3(impl2)
	fmt.Println(impl3.Render())
}
