package internal

import (
	"context"

	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
)

type Sms interface {
	Send(ctx context.Context, deliver *deliver.Sms) (id string, status kernel.Status, err error)
}
