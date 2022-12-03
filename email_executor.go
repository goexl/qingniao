package qingniao

import (
	"context"
)

type emailExecutor interface {
	send(ctx context.Context, deliver *emailDeliver) (id string, err error)
}
