package qingniao

type chuangcacheCode struct {
	success        int
	badRequest     int
	unauthorized   int
	userNotfound   int
	appNotfound    int
	failed         int
	noBalance      int
	timestampError int
	tokenError     int
	mobileInvalid  int
}

func newChuangcacheCode() chuangcacheCode {
	return chuangcacheCode{
		success:        1000,
		badRequest:     2001,
		unauthorized:   2003,
		userNotfound:   2005,
		appNotfound:    2009,
		failed:         2019,
		noBalance:      2011,
		timestampError: 2012,
		tokenError:     2013,
		mobileInvalid:  2015,
	}
}
