package facade

type Facade struct {
	Dep1 *Dependency1
	Dep2 *Dependency2
	Dep3 *Dependency3
}

func NewFacade() *Facade {
	return &Facade{
		Dep1: NewDependency1(),
		Dep2: NewDependency2(),
		Dep3: NewDependency3(),
	}
}

func (f *Facade) Execute(dependency int) {
	switch dependency {
	case 1:
		f.Dep1.Execute()
	case 2:
		f.Dep2.Execute()
	case 3:
		f.Dep3.Execute()
	}
}
