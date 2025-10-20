package adapter

import "fmt"

/*
	Adapter is a pattern that allows incompatible interfaces to work together,
	"adapting" one interface to another expected by the client, without changing the underlying implementation.

	PROS:
	- Enables compatibility - Makes two incompatible interfaces work together without modifying them.
	- Decouples code - Client code depends on the interface, not on concrete implementations.
	- Reusability - Lets you reuse existing code or libraries with a new interface.
	- Single Responsibility - Keeps adaptation logic separate from the original code.
	- Supports polymorphism - Adapters can be passed wherever the target interface is expected.
	CONS:
	- Extra layer - Adds a wrapper, which can increase complexity.
	- Potential overhead - Slight performance or memory overhead from indirection.
	- Doesn't fix fundamental design flaws - Only makes interfaces compatible; underlying logic may still be cumbersome.
	- Can proliferate - Too many adapters can clutter the codebase if overused.

	WHEN TO USE:
	- You need to integrate a third-party library that has an incompatible interface.
	- You want to reuse legacy code without modifying it.
	- You want to decouple your code from implementation details.
	- You need polymorphism with incompatible types.
	- You want to switch between multiple implementatios transparently.

*/

type TargetSourceAdapter struct {
	*Target
	Source *Source
}

func NewTargetSourceAdapter() *TargetSourceAdapter {
	return &TargetSourceAdapter{
		Target: NewTarget(),
		Source: NewSource(),
	}
}

func (a *TargetSourceAdapter) Execute(arg string) {
	fmt.Println("Adapting Target to use Source's method")
	a.Source.Execute(arg)
}
