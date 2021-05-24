package una

import (
	`context`
)

type unaInternal interface {
	send(ctx context.Context, content string, options *options) (id string, err error)
}
