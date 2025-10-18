package singleton

import (
	"sync"
)

/*
	A Singleton is a class that only gets instantiated once

	PROS:
	- Ensures a single instance - Guarantees that only one instance of a given type exists.
	- Centralized access point - Simplifies global coordination and reuse of the same object.
	- Lazy initialization possible - Can delay creation until first needed (sync.Once).
	- Saves resources - Prevents duplicate instances of expensive objects.
	- Encapsulates initialization - Keeps setup logic hidden from consumers.
	CONS:
	- Global state - Makes behavior harder to predict and reason about.
	- Testing difficulty - Hard to mock or replace in tests because it's globally accessible.
	- Hidden dependencies - Functions may rely on it implicitly, breaking modularity.
	- Tight coupling - Code becomes dependent on a single concrete instance.
	- Potential concurrency issues - Unsafe if implemented without proper synchronization.

	WHEN TO USE:
	- When you need exactly one instance of something across the whole app.
	- The object is stateless or immutable, so shared access is safe.
	- The resource is expensive to initialize and must be reused globally.
	- You want lazy initialization that's thread safe.
*/

type Singleton struct {
	property string
	list     []string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}

func (s *Singleton) SetProperty(str string) {
	s.property = str
}

func (s *Singleton) GetProperty() string {
	return s.property
}

func (s *Singleton) AddItem(item string) {
	s.list = append(s.list, item)
}

func (s *Singleton) GetListLen() int {
	return len(s.list)
}
