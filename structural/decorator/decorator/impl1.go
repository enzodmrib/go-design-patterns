package decorator

import "fmt"

type Impl1 struct {
	Wrapped Decorator
}

func NewImpl1(d Decorator) *Impl1 {
	return &Impl1{
		Wrapped: d,
	}
}

func (d *Impl1) Render() string {
	return fmt.Sprintf("Impl1: %s", d.Wrapped.Render())
}
