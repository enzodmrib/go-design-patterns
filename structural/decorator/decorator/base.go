package decorator

type Base struct {
	Content string
}

func NewBase(content string) *Base {
	return &Base{Content: content}
}

func (b *Base) Render() string {
	return b.Content
}
