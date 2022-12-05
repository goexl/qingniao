package qingniao

import (
	"context"
)

type smsExecutor interface {
	send(ctx context.Context, deliver *smsDeliver) (id string, status Status, err error)
}
