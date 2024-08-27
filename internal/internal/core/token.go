package core

import (
	"time"
)

type Token struct {
	Code      string
	ExpiresIn time.Time
}

func (t *Token) Validate() bool {
	return time.Now().Before(t.ExpiresIn.Add(5 * time.Minute))
}
