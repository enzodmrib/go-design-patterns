package main

import "github.com/enzodmrib/go-design-patterns/behavioral/strategy/strategy"

func main() {
	impl := strategy.NewImpl()

	impl.SetStrategy(strategy.NewOption1())
	impl.Execute()

	impl.SetStrategy(strategy.NewOption2())
	impl.Execute()
}
