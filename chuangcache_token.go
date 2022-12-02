package qingniao

import (
	"time"
)

type chuangcacheToken struct {
	token     string
	expiresIn time.Time
}

func (ct *chuangcacheToken) validate() bool {
	return time.Now().After(ct.expiresIn.Add(5 * time.Minute))
}
