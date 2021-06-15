package una

import (
	`github.com/storezhang/pangu`
	_ `github.com/storezhang/pangu-http`
)

func init() {
	if err := pangu.New().Provides(
		NewChuangcache,
		NewEmail,
		New,
	); nil != err {
		panic(err)
	}
}
