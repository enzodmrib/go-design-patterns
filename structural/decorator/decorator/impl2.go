package decorator

import "fmt"

type Impl2 struct {
	Wrapped Decorator
}

func NewImpl2(d Decorator) *Impl2 {
	return &Impl2{
		Wrapped: d,
	}
}

func (d *Impl2) Render() string {
	return fmt.Sprintf("Impl2: %s", d.Wrapped.Render())
}
