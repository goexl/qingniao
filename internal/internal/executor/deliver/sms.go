package deliver

import (
	"github.com/goexl/qingniao/internal/internal/constant"
)

type Sms struct {
	Key     string           `validate:"required"`
	Mobiles []string         `validate:"required,dive,e164"`
	Content string           `validate:"required,max=536"`
	Type    constant.SmsType `validate:"oneof=1 2 3"`
}
