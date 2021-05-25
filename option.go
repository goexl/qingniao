package una

type option interface {
	apply(options *options)
}

// NewOptions 创建空选项列表
func NewOptions(opts ...option) []option {
	return opts
}
