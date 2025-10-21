package main

import (
	"github.com/enzodmrib/go-design-patterns/structural/facade/facade"
)

func main() {
	facade := facade.NewFacade()

	facade.Execute(1)
	facade.Execute(2)
	facade.Execute(3)
}
