package qingniao

import (
	"context"
)

type smsDeliver struct {
	AppKey  string   `validate:"required"`
	Mobiles []string `validate:"required,unique"`
	Content string   `validate:"required,max=536"`
	Type    smsType  `validate:"oneof=1 2"`

	executor smsExecutor
}

func newSmsDeliver(appKey string, mobiles []string, content string, executor smsExecutor) *smsDeliver {
	return &smsDeliver{
		AppKey:   appKey,
		Mobiles:  mobiles,
		Content:  content,
		executor: executor,
	}
}

func (sd *smsDeliver) To(mobiles ...string) *smsDeliver {
	sd.Mobiles = append(sd.Mobiles, mobiles...)

	return sd
}

func (sd *smsDeliver) Code() *smsDeliver {
	sd.Type = smsTypeCode

	return sd
}

func (sd *smsDeliver) Notify() *smsDeliver {
	sd.Type = smsTypeNotify

	return sd
}

func (sd *smsDeliver) Advertising() *smsDeliver {
	sd.Type = smsTypeAdvertising

	return sd
}

func (sd *smsDeliver) Send(ctx context.Context) (string, Status, error) {
	return sd.executor.send(ctx, sd)
}
