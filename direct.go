package qingniao

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/goexl/simaqian"
	"github.com/goexl/xiren"
	"github.com/jordan-wright/email"
)

var _ emailExecutor = (*direct)(nil)

type direct struct {
	host     string
	port     int
	username string
	password string
	identity string
	poolSize int

	pool   *email.Pool
	logger simaqian.Logger
}

func newDirect(
	host string, port int,
	username string, password string, identity string,
	poolSize int,
	logger simaqian.Logger,
) *direct {
	return &direct{
		host:     host,
		port:     port,
		username: username,
		password: password,
		identity: identity,
		poolSize: poolSize,

		logger: logger,
	}
}

func (d *direct) send(_ context.Context, deliver *emailDeliver) (id string, err error) {
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
	em.From = deliver.from
	if "" != em.From {
		em.From = fmt.Sprintf("%s@%s", d.username, d.host)
	}
	em.To = deliver.to
	em.Bcc = deliver.bcc
	em.Cc = deliver.cc
	em.Subject = deliver.subject
	switch deliver.typ {
	case emailTypeHtml:
		em.HTML = []byte(deliver.content)
	case emailTypePlain:
		em.Text = []byte(deliver.content)
	default:
		em.HTML = []byte(deliver.content)
	}
	err = d.pool.Send(em, deliver.timeout)

	return
}
