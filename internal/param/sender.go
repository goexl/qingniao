package param

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type Sender struct {
	Http   *http.Client
	Logger log.Logger
}

func NewSender() *Sender {
	return &Sender{
		Http:   http.New().Build(),
		Logger: log.New().Apply(),
	}
}
