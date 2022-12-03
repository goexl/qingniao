package qingniao

import (
	"context"
)

type smsDeliver struct {
	mobiles []string `validate:"required,dive,mobile"`
	content string   `validate:"required,max=536"`
	typ     smsType  `validate:"oneof=1 2 3"`

	executor smsExecutor
}

func newSmsDeliver(mobile string, content string, executor smsExecutor) *smsDeliver {
	return &smsDeliver{
		mobiles:  []string{mobile},
		content:  content,
		executor: executor,
	}
}

func (sd *smsDeliver) To(mobiles ...string) *smsDeliver {
	sd.mobiles = append(sd.mobiles, mobiles...)

	return sd
}

func (sd *smsDeliver) Send(ctx context.Context) (string, error) {
	return sd.executor.send(ctx, sd)
}
