package chuangcache

type Code struct {
	Success        int
	BadRequest     int
	Unauthorized   int
	UserNotfound   int
	AppNotfound    int
	Failed         int
	NoBalance      int
	TimestampError int
	TokenError     int
	MobileInvalid  int
}

func NewCode() Code {
	return Code{
		Success:        1000,
		BadRequest:     2001,
		Unauthorized:   2003,
		UserNotfound:   2005,
		AppNotfound:    2009,
		Failed:         2019,
		NoBalance:      2011,
		TimestampError: 2012,
		TokenError:     2013,
		MobileInvalid:  2015,
	}
}
