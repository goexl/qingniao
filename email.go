package qingniao

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// Email 邮件
type Email struct {
	pool     *email.Pool
	username string
	host     string
}

func newEmail(
	host string, port int,
	username string, password string, identity string,
	poolSize int,
) (em *Email, err error) {
	em = new(Email)
	em.username = username
	em.host = host

	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth(identity, username, password, host)
	em.pool, err = email.NewPool(addr, poolSize, auth)

	return
}

func (e *Email) NewDeliver(subject string, content string) *emailDeliver {
	return newEmailDeliver(fmt.Sprintf("%s@%s", e.username, e.host), subject, content, e.pool)
}
