package adapter

type Target struct{}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) Execute(arg string) {
	panic("execute() method must be implemented")
}
