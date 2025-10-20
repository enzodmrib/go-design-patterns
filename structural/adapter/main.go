package main

import (
	"github.com/enzodmrib/go-design-patterns/structural/adapter/adapter"
)

func main() {
	adapter := adapter.NewTargetSourceAdapter()

	adapter.Execute("Hello!!")
}
