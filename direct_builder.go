package qingniao

import (
	"runtime"

	"github.com/goexl/log"
)

type directBuilder struct {
	host     string
	port     int
	username string
	password string
	poolSize int
	identity string

	logger log.Logger
}

func newDirectBuilder(host string, port int, logger log.Logger) *directBuilder {
	return &directBuilder{
		host:     host,
		port:     port,
		poolSize: runtime.NumCPU() + 1,

		logger: logger,
	}
}

func (eb *directBuilder) PoolSize(size int) *directBuilder {
	eb.poolSize = size

	return eb
}

func (eb *directBuilder) Identity(identity string) *directBuilder {
	eb.identity = identity

	return eb
}

func (eb *directBuilder) Build() *Email {
	return newEmail(newDirect(eb.host, eb.port, eb.username, eb.password, eb.identity, eb.poolSize, eb.logger))
}
