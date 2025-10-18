package main

import (
	"fmt"

	"github.com/enzodmrib/go-design-patterns/creational/singleton/singleton"
)

func main() {
	instance1 := singleton.GetInstance()

	instance1.SetProperty("One")

	instance1.AddItem("One")

	instance2 := singleton.GetInstance()

	instance2.SetProperty("Two")

	instance2.AddItem("Two")

	fmt.Printf("Instance1 property: %v\n", instance1.GetProperty())
	fmt.Printf("Instance1 list length: %v\n", instance1.GetListLen())

	fmt.Printf("Instance2 property: %v\n", instance2.GetProperty())
	fmt.Printf("Instance2 list length: %v\n", instance2.GetListLen())

	/*
		Output:
		Instance1 property: Two
		Instance1 list length: 2
		Instance2 property: Two
		Instance2 list length: 2

		Instances are the same!
	*/
}
