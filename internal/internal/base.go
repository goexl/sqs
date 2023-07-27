package internal

type Base struct {
	Label string
	Queue string
	Url   string
}

func NewBase() *Base {
	return &Base{
		Label: DefaultLabel,
	}
}
