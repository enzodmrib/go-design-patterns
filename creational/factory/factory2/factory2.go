package factory2

import "fmt"

type Implementable interface {
	Connect() string
}

var factories = map[int]func() Implementable{}

func RegisterFactoryFn(key int, function func() Implementable) {
	factories[key] = function
}

func init() {
	RegisterFactoryFn(1, func() Implementable { return &Impl1{} })
	RegisterFactoryFn(2, func() Implementable { return &Impl2{} })
}

type Impl1 struct{}

func (i *Impl1) Connect() string {
	return "Connected to implementation 1"
}

type Impl2 struct{}

func (i *Impl2) Connect() string {
	return "Connected to implementation 2"
}

func NewFactoryImpl(key int) (Implementable, error) {
	f, ok := factories[key]
	if !ok {
		return nil, fmt.Errorf("unknown implementation: %d", key)
	}

	return f(), nil
}
