package una

import (
	`context`
)

type unaTemplate struct {
	una Una
}

func (t *unaTemplate) Send(ctx context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	id, err = t.una.Send(ctx, content, opts...)

	return
}
