package main

import (
	"fmt"

	"github.com/enzodmrib/go-design-patterns/behavioral/iterator/iterator"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	list := []struct {
		Name string
		Age  int
	}{
		{
			Name: "Enzo",
			Age:  1,
		},
		{
			Name: "Other",
			Age:  2,
		},
	}

	iterator := iterator.NewIterator(list)

	for iterator.Next() {
		fmt.Println(iterator.Get())
	}

}
