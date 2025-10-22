package strategy

import "fmt"

type Option3 struct {
}

func NewOption3() *Option3 {
	return &Option3{}
}

func (o *Option3) Execute() {
	fmt.Println("Executing option 3")
}
