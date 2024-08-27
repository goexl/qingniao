package deliver

import (
	"time"

	"github.com/goexl/qingniao/internal/internal/constant"
)

type Email struct {
	Type    constant.EmailType `validate:"required,oneof=html plain"`
	Subject string             `validate:"required"`
	Content string             `validate:"required"`
	From    string             `validate:"required"`
	To      []string           `validate:"required,dive,email"`
	Cc      []string           `validate:"omitempty,dive,email"`
	Bcc     []string           `validate:"omitempty,dive,email"`
	Timeout time.Duration      `validate:"required,min=1"`
}
