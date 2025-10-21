package facade

import "fmt"

type Dependency1 struct{}

func NewDependency1() *Dependency1 {
	return &Dependency1{}
}

func (d *Dependency1) Execute() {
	fmt.Println("Executing from 1")
}

type Dependency2 struct{}

func NewDependency2() *Dependency2 {
	return &Dependency2{}
}

func (d *Dependency2) Execute() {
	fmt.Println("Executing from 2")
}

type Dependency3 struct{}

func NewDependency3() *Dependency3 {
	return &Dependency3{}
}

func (d *Dependency3) Execute() {
	fmt.Println("Executing from 3")
}
