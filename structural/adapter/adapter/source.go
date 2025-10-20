package adapter

import "fmt"

type Source struct{}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Execute(arg string) {
	fmt.Println("Executing from source: ", arg)
}
