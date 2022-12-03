package qingniao

import (
	"time"
)

type token struct {
	token     string
	expiresIn time.Time
}

func (t *token) validate() bool {
	return time.Now().After(t.expiresIn.Add(5 * time.Minute))
}
