package constant

const (
	ExecutorUnknown Executor = iota
	ExecutorDirect
	ExecutorServerChain
	ExecutorChuangcache
)

type Executor uint8
