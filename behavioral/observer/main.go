package main

import "github.com/enzodmrib/go-design-patterns/behavioral/observer/observer"

func main() {
	s := observer.NewConcreteSubject()
	o1 := observer.NewConcreateObserver("1")
	o2 := observer.NewConcreateObserver("2")

	s.AddObserver(o1)
	s.AddObserver(o2)

	s.RemoveObserver(o1)

	s.Receive()
}
