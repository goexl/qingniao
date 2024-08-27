package internal

import (
	"context"

	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
)

type Wechat interface {
	Send(ctx context.Context, deliver *deliver.Wechat) (id string, status kernel.Status, err error)
}
