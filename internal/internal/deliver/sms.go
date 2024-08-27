package deliver

import (
	"context"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
	"github.com/goexl/xiren"
)

type Sms struct {
	base

	key     string
	mobiles []string
	content string
	typ     constant.SmsType

	executors map[string]internal.Sms
}

func NewSms(mobile string, content string, executors map[string]internal.Sms) (sms *Sms) {
	return &Sms{
		base: newBase(),

		mobiles: []string{mobile},
		content: content,

		executors: executors,
	}
}

func (s *Sms) AppKey(key string) (sms *Sms) {
	s.key = key
	sms = s

	return
}

func (s *Sms) To(mobile string, mobiles ...string) (sms *Sms) {
	s.mobiles = append(s.mobiles, mobile)
	s.mobiles = append(s.mobiles, mobiles...)
	sms = s

	return
}

func (s *Sms) Code() (sms *Sms) {
	s.typ = constant.SmsTypeCode
	sms = s

	return
}

func (s *Sms) Notify() (sms *Sms) {
	s.typ = constant.SmsTypeNotify
	sms = s

	return
}

func (s *Sms) Advertising() (sms *Sms) {
	s.typ = constant.SmsTypeAdvertising
	sms = s

	return
}

func (s *Sms) Send(ctx context.Context) (id string, status kernel.Status, err error) {
	message := new(deliver.Sms)
	message.Key = s.key
	message.Mobiles = s.mobiles
	message.Content = s.content
	message.Type = s.typ

	label := s.base.params.Label
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, ok := s.executors[label]; !ok {
		err = exception.New().Message("没有找到短信执行器").
			Field(field.New("executors", s.executors)).Field(field.New("label", label)).
			Build()
	} else {
		id, status, err = executor.Send(ctx, message)
	}

	return
}
