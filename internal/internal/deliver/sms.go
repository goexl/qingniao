package deliver

import (
	"context"

	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
	"github.com/goexl/xiren"
)

type Sms struct {
	key     string
	mobiles []string
	content string
	typ     constant.SmsType

	picker  *picker[internal.Sms]
	current constant.Executor
}

func NewSms(mobile string, content string, executors map[constant.Executor]internal.Sms) (sms *Sms) {
	return &Sms{
		mobiles: []string{mobile},
		content: content,

		picker: newPicker(executors),
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
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, pe := s.picker.pick(s.current, "短信"); nil != pe {
		err = pe
	} else {
		id, status, err = executor.Send(ctx, message)
	}

	return
}

func (s *Sms) Chuangcache() (sms *Sms) {
	s.current = constant.ExecutorChuangcache
	sms = s

	return
}
