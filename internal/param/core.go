package param

type Core struct {
	Label string
}

func NewCore() *Core {
	return &Core{
		Label: "default",
	}
}
