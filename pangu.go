package una

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Sets(
		NewChuangcacheSms,
		New,
	); nil != err {
		panic(err)
	}
}
