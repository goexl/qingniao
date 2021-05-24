package una

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Sets(
		NewChuangcache,
		NewEmail,
		New,
	); nil != err {
		panic(err)
	}
}
