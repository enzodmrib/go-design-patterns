package main

import (
	"fmt"

	"github.com/enzodmrib/go-design-patterns/creational/builder/builder"
)

func main() {
	b := builder.NewBuilder()

	obj := b.Property1("non-default").Property2(2).Property3(false).Build()

	fmt.Println(obj)
}
