package qingniao

import (
	"runtime"
)

type emailBuilder struct {
	host     string
	port     int
	username string
	password string
	poolSize int
	identity string
}

func newEmailBuilder(host string, port int) *emailBuilder {
	return &emailBuilder{
		host:     host,
		port:     port,
		poolSize: runtime.NumCPU() + 1,
	}
}

func (eb *emailBuilder) PoolSize(size int) *emailBuilder {
	eb.poolSize = size

	return eb
}

func (eb *emailBuilder) Identity(identity string) *emailBuilder {
	eb.identity = identity

	return eb
}

func (eb *emailBuilder) Build() (*Email, error) {
	return newEmail(eb.host, eb.port, eb.username, eb.password, eb.identity, eb.poolSize)
}
