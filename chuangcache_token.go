package una

import (
	`time`
)

type chuangcacheToken struct {
	token     string
	expiresIn time.Time
}

func (ct *chuangcacheToken) validate() (token string, validate bool) {
	return ct.token, time.Now().After(ct.expiresIn.Add(5 * time.Minute))
}
