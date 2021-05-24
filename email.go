package una

import (
	`context`
	`fmt`
	`net/smtp`
	`sync`
	`time`

	`github.com/jordan-wright/email`
	`github.com/storezhang/validatorx`
)

// Email 邮件通知
type Email struct {
	poolCache sync.Map
}

// NewEmail 创建普通邮件
func NewEmail() *Email {
	return &Email{
		poolCache: sync.Map{},
	}
}

func (e *Email) Send(_ context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	if err = validatorx.Validate(options.mail); nil != err {
		return
	}

	var pool *email.Pool
	if pool, err = e.getPool(options); nil != err {
		return
	}

	em := email.NewEmail()
	em.From = "Jordan Wright <test@gmail.com>"
	em.To = options.mail.to
	em.Bcc = options.mail.bcc
	em.Cc = options.mail.cc
	em.Subject = options.mail.subject
	switch options.mail.emailType {
	case EmailTypeHtml:
		em.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	case EmailTypePlain:
		em.Text = []byte("Text Body is, of course, supported!")
	default:
		em.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	}
	err = pool.Send(em, 10*time.Second)

	return
}

func (e *Email) getPool(options *options) (pool *email.Pool, err error) {
	var (
		cache interface{}
		ok    bool
	)

	key := options.mail.key()
	if cache, ok = e.poolCache.Load(key); ok {
		pool = cache.(*email.Pool)

		return
	}

	if pool, err = email.NewPool(
		fmt.Sprintf("%s:%d", options.mail.host, options.mail.port),
		options.poolSize,
		smtp.PlainAuth("", options.mail.username, options.mail.password, options.mail.host),
	); nil != err {
		return
	}
	e.poolCache.Store(key, pool)

	return
}
