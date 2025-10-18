package factory

import "fmt"

/*
	The factory pattern provides an interface for creating objects but lets subclasses decide which specific class to instantiate

	PROS:
	- Centralizes object creation - You don't repeat the same if/switch logic everywhere.
	- Hides construction details - The client doesn't need to know how the object is built.
	- Improves extensibility - Adding new types doesn't require changing client code.
	- Promotes dependency inversion - Client depends on an interface, not a concrete type.
	- Supports runtime selection - Choose which implementation to use dynamically.
	CONS:
	- Adds indirection - You have to look inside the factory to see what's really created.
	- Simple cases get overcomplicated - For one or two types, a plain NewX() function is clearer.
	- Switch logic still exists - It's centralized but not eliminated - still must maintain it.
	- Less transparent for debugging - Sometimes hides which conecrete type is used under the hood.

	WHEN TO USE:
	- When you have multimple implementations of an interface (e.g., Database, Storage, Cache).
	- When creation depends on runtime data (e.g., confi, environment, user input).
	- When you want plugins or drivers that can register themselves dynamically.
*/

type Implementable interface {
	Connect() string
}

type Impl1 struct{}

func (i *Impl1) Connect() string {
	return "Connected to implementation 1"
}

type Impl2 struct{}

func (i *Impl2) Connect() string {
	return "Connected to implementation 2"
}

func NewFactoryImpl(impl int) (Implementable, error) {
	switch impl {
	case 1:
		return &Impl1{}, nil
	case 2:
		return &Impl2{}, nil
	default:
		return nil, fmt.Errorf("unknown implementation: %d", impl)
	}
}
