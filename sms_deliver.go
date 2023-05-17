package qingniao

import (
	"context"
)

type (
	smsDeliver struct {
		key     any
		mobile  []string
		content string
		typ     smsType
		executor smsExecutor
	}

	smsDeliverInternal struct {
		Key     any      `validate:"required"`
		Mobiles []string `validate:"required,unique"`
		Content string   `validate:"required,max=536"`
		Type    smsType  `validate:"oneof=1 2 3"`

		executor smsExecutor
	}
)

func newSmsDeliver(mobiles []string, content string, executor smsExecutor) *smsDeliver {
	return &smsDeliver{
		mobile:   mobiles,
		content:  content,
		executor: executor,
	}
}

func (sd *smsDeliver) AppKey(key string) *smsDeliver {
	sd.key = key

	return sd
}

func (sd *smsDeliver) To(mobiles ...string) *smsDeliver {
	sd.mobile = append(sd.mobile, mobiles...)

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
	return sd.executor.send(ctx, &smsDeliverInternal{
		Key:      sd.key,
		Mobiles:  sd.mobile,
		Content:  sd.content,
		Type:     sd.typ,
		executor: sd.executor,
	})
}
