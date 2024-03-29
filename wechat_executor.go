package qingniao

import (
	"context"
)

type wechatExecutor interface {
	send(ctx context.Context, deliver *wechatDeliver) (id string, status Status, err error)
}
