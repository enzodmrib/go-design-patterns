package strategy

type Impl struct {
	strategy Strategy
}

func NewImpl() *Impl {
	return &Impl{}
}

func (i *Impl) SetStrategy(strategy Strategy) {
	i.strategy = strategy
}

func (i *Impl) Execute() {
	i.strategy.Execute()
}
