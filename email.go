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
	validate  validatorx.Validate
	poolCache sync.Map
}

// NewEmail 创建普通邮件
func NewEmail(validate validatorx.Validate) *Email {
	return &Email{
		validate:  validate,
		poolCache: sync.Map{},
	}
}

func (e *Email) Send(_ context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	if err = e.validate.Struct(options.email); nil != err {
		return
	}

	var pool *email.Pool
	if pool, err = e.getPool(options); nil != err {
		return
	}

	em := email.NewEmail()
	em.From = options.email.from
	em.To = options.email.to
	em.Bcc = options.email.bcc
	em.Cc = options.email.cc
	em.Subject = options.email.subject
	switch options.email.emailType {
	case EmailTypeHtml:
		em.HTML = []byte(content)
	case EmailTypePlain:
		em.Text = []byte(content)
	default:
		em.HTML = []byte(content)
	}
	err = pool.Send(em, 10*time.Second)

	return
}

func (e *Email) getPool(options *options) (pool *email.Pool, err error) {
	var (
		cache interface{}
		ok    bool
	)

	key := options.email.key()
	if cache, ok = e.poolCache.Load(key); ok {
		pool = cache.(*email.Pool)

		return
	}

	if pool, err = email.NewPool(
		fmt.Sprintf("%s:%d", options.email.host, options.email.port),
		options.poolSize,
		smtp.PlainAuth("", options.email.username, options.email.password, options.email.host),
	); nil != err {
		return
	}
	e.poolCache.Store(key, pool)

	return
}
