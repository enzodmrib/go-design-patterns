package strategy

import "fmt"

type Option2 struct {
}

func NewOption2() *Option2 {
	return &Option2{}
}

func (o *Option2) Execute() {
	fmt.Println("Executing option 2")
}
