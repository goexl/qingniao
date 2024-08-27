package builder

import (
	"github.com/goexl/qingniao/internal/param"
)

type base[T any] struct {
	params *param.Base
	t      *T
}

func newBase[T any](t *T) *base[T] {
	return &base[T]{
		params: param.NewBase(),
		t:      t,
	}
}

func (b *base[T]) Label(label string) (t *T) {
	b.params.Label = label
	t = b.t

	return
}
