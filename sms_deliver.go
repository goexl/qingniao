package qingniao

import (
	"context"
)

type smsDeliver struct {
	template string   `validate:"required"`
	mobiles  []string `validate:"required,dive,mobile"`
	content  string   `validate:"required,max=536"`
	typ      smsType  `validate:"oneof=1 2"`

	executor smsExecutor
}

func newSmsDeliver(template string, mobiles []string, content string, executor smsExecutor) *smsDeliver {
	return &smsDeliver{
		template: template,
		mobiles:  mobiles,
		content:  content,
		executor: executor,
	}
}

func (sd *smsDeliver) To(mobiles ...string) *smsDeliver {
	sd.mobiles = append(sd.mobiles, mobiles...)

	return sd
}

func (sd *smsDeliver) Code() *smsDeliver {
	sd.typ = smsTypeCode

	return sd
}

func (sd *smsDeliver) Notify() *smsDeliver {
	sd.typ = smsTypeNotify

	return sd
}

func (sd *smsDeliver) Advertising() *smsDeliver {
	sd.typ = smsTypeAdvertising

	return sd
}

func (sd *smsDeliver) Send(ctx context.Context) (string, Status, error) {
	return sd.executor.send(ctx, sd)
}
