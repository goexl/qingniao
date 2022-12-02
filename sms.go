package qingniao

import (
	"github.com/go-resty/resty/v2"
)

// Sms 短信
type Sms struct {
	http *resty.Client
}

func newSms(client *resty.Client) *Sms {
	return &Sms{
		http: client,
	}
}

func (s *Sms) To(phone string) *Sms {

	return s
}
