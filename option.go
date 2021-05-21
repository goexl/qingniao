package una

type option interface {
	apply(options *options)
}
