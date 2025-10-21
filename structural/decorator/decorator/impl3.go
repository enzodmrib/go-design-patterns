package decorator

import "fmt"

type Impl3 struct {
	Wrapped Decorator
}

func NewImpl3(d Decorator) *Impl3 {
	return &Impl3{
		Wrapped: d,
	}
}

func (d *Impl3) Render() string {
	return fmt.Sprintf("Impl3: %s", d.Wrapped.Render())
}
