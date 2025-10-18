package main

import (
	"fmt"

	"github.com/enzodmrib/go-design-patterns/creational/factory/factory"
	"github.com/enzodmrib/go-design-patterns/creational/factory/factory2"
)

func main() {
	impl, err := factory.NewFactoryImpl(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(impl.Connect())

	impl2, err := factory2.NewFactoryImpl(2)
	if err != nil {
		panic(err)
	}

	fmt.Println(impl2.Connect())
}
