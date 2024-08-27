package param

type Base struct {
	Label string
}

func NewBase() *Base {
	return &Base{
		Label: "default",
	}
}
