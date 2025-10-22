package observer

import "fmt"

type Observer interface {
	Update()
}

type ConcreteObserver struct {
	name string
}

func NewConcreateObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{
		name: name,
	}
}

func (o *ConcreteObserver) Update() {
	fmt.Printf("Observer %s: Notified\n", o.name)
}
