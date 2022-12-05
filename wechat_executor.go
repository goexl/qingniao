package qingniao

import (
	"context"
)

type wechatExecutor interface {
	send(ctx context.Context, deliver *smsDeliver) (id string, status Status, err error)
}
