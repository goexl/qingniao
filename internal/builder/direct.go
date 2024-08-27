package builder

import (
	"runtime"

	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/internal/executor"
)

type Direct struct {
	host     string
	port     int
	username string
	password string
	poolSize int
	identity string

	logger log.Logger
	email  *Email
}

func NewDirect(host string, port int, logger log.Logger, email *Email) (direct *Direct) {
	return &Direct{
		host:     host,
		port:     port,
		poolSize: runtime.NumCPU() + 1,

		logger: logger,
		email:  email,
	}
}

func (d *Direct) PoolSize(size int) (direct *Direct) {
	d.poolSize = size
	direct = d

	return
}

func (d *Direct) Identity(identity string) (direct *Direct) {
	d.identity = identity
	direct = d

	return
}

func (d *Direct) Build() (email *Email) {
	d.email.direct = executor.NewDirect(d.host, d.port, d.username, d.password, d.identity, d.poolSize, d.logger)
	email = d.email

	return
}
