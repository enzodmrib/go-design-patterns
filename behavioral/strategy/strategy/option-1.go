package strategy

import "fmt"

type Option1 struct {
}

func NewOption1() *Option1 {
	return &Option1{}
}

func (o *Option1) Execute() {
	fmt.Println("Executing option 1")
}
