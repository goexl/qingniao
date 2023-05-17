package qingniao

import (
	"context"
)

type smsExecutor interface {
	send(ctx context.Context, deliver *smsDeliverInternal) (id string, status Status, err error)
}
