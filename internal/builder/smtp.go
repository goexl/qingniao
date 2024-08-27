package builder

import (
	"runtime"

	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/internal/executor"
)

type Smtp struct {
	host     string
	port     int
	username string
	password string
	poolSize int
	identity string

	logger log.Logger
	email  *Email
}

func NewSmtp(host string, port int, logger log.Logger, email *Email) (direct *Smtp) {
	return &Smtp{
		host:     host,
		port:     port,
		poolSize: runtime.NumCPU() + 1,

		logger: logger,
		email:  email,
	}
}

func (s *Smtp) Pool(size int) (direct *Smtp) {
	s.poolSize = size
	direct = s

	return
}

func (s *Smtp) Identity(identity string) (direct *Smtp) {
	s.identity = identity
	direct = s

	return
}

func (s *Smtp) Username(username string) (direct *Smtp) {
	s.username = username
	direct = s

	return
}

func (s *Smtp) Password(password string) (direct *Smtp) {
	s.password = password
	direct = s

	return
}

func (s *Smtp) Build() (email *Email) {
	s.email.direct = executor.NewSmtp(s.host, s.port, s.username, s.password, s.identity, s.poolSize, s.logger)
	email = s.email

	return
}
