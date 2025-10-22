package observer

import (
	"fmt"
	"slices"
)

type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type ConcreteSubject struct {
	observers []Observer
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (s *ConcreteSubject) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	s.observers = slices.DeleteFunc(s.observers, func(o Observer) bool {
		return o == observer
	})
}

func (s *ConcreteSubject) NotifyObservers() {
	fmt.Println("Notifying subscribers")
	for _, o := range s.observers {
		o.Update()
	}
}

func (s *ConcreteSubject) Receive() {
	fmt.Println("Received")
	s.NotifyObservers()
}
