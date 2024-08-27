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

var _ internal.Email = (*Direct)(nil)

type Direct struct {
	host     string
	port     int
	username string
	password string
	identity string
	poolSize int

	pool   *email.Pool
	logger log.Logger
}

func NewDirect(
	host string, port int,
	username string, password string, identity string,
	poolSize int,
	logger log.Logger,
) *Direct {
	return &Direct{
		host:     host,
		port:     port,
		username: username,
		password: password,
		identity: identity,
		poolSize: poolSize,

		logger: logger,
	}
}

func (d *Direct) Send(_ context.Context, deliver *deliver.Email) (id string, err error) {
	if err = xiren.Struct(deliver); nil != err {
		return
	}

	if nil == d.pool {
		addr := fmt.Sprintf("%s:%d", d.host, d.port)
		auth := smtp.PlainAuth(d.identity, d.username, d.password, d.host)
		d.pool, err = email.NewPool(addr, d.poolSize, auth)
	}
	if nil != err {
		return
	}

	em := email.NewEmail()
	em.From = deliver.From
	if "" != em.From {
		em.From = fmt.Sprintf("%s@%s", d.username, d.host)
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
	err = d.pool.Send(em, deliver.Timeout)

	return
}
