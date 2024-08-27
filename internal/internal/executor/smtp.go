package executor

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/xiren"
	"github.com/jordan-wright/email"
)

var _ internal.Email = (*Smtp)(nil)

type Smtp struct {
	host     string
	port     int
	username string
	password string
	identity string
	poolSize int

	pool   *email.Pool
	logger log.Logger
}

func NewSmtp(
	host string, port int,
	username string, password string, identity string,
	poolSize int,
	logger log.Logger,
) *Smtp {
	return &Smtp{
		host:     host,
		port:     port,
		username: username,
		password: password,
		identity: identity,
		poolSize: poolSize,

		logger: logger,
	}
}

func (s *Smtp) Send(_ context.Context, deliver *deliver.Email) (id string, err error) {
	if err = xiren.Struct(deliver); nil != err {
		return
	}

	if nil == s.pool {
		addr := fmt.Sprintf("%s:%d", s.host, s.port)
		auth := smtp.PlainAuth(s.identity, s.username, s.password, s.host)
		s.pool, err = email.NewPool(addr, s.poolSize, auth)
	}
	if nil != err {
		return
	}

	em := email.NewEmail()
	em.From = deliver.From
	if "" != em.From {
		em.From = fmt.Sprintf("%s@%s", s.username, s.host)
	}
	em.To = deliver.To
	em.Bcc = deliver.Bcc
	em.Cc = deliver.Cc
	em.Subject = deliver.Subject
	switch deliver.Type {
	case constant.EmailTypeHtml:
		em.HTML = []byte(deliver.Content)
	case constant.EmailTypePlain:
		em.Text = []byte(deliver.Content)
	default:
		em.HTML = []byte(deliver.Content)
	}
	err = s.pool.Send(em, deliver.Timeout)

	return
}
