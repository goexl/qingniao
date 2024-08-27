package internal

import (
	"context"

	"github.com/goexl/qingniao/internal/internal/executor/deliver"
)

type Email interface {
	Send(ctx context.Context, deliver *deliver.Email) (id string, err error)
}
