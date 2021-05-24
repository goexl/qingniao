package una

import (
	`context`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/validatorx`
)

// Una 通知接口
type Una interface {
	// Send 发送消息
	Send(ctx context.Context, content string, opts ...option) (id string, err error)
}

// New 创建适配器
func New(tpe Type, validate *validatorx.Validate, resty *resty.Request) (una Una, err error) {
	if err = validate.Var(tpe, "required,oneof=email chuangcache"); nil != err {
		return
	}

	switch tpe {
	case TypeChuangcache:
		una = NewChuangcacheSms(validate, resty)
	case TypeEmail:
		una = NewEmail(validate)
	}

	return
}
