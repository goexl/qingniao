package deliver

import (
	"github.com/goexl/qingniao/internal/param"
)

type base struct {
	params *param.Core
}

func newBase() base {
	return base{
		params: param.NewCore(),
	}
}

func (b *base) Label(label string) (core *base) {
	b.params.Label = label
	core = b

	return
}
