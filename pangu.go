package una

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Provides(
		NewChuangcache,
		NewEmail,
		New,
	); nil != err {
		panic(err)
	}
}
