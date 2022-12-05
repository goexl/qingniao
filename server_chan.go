package qingniao

type serverChan struct {
	token string
}

func newServerChan(token string) *serverChan {
	return &serverChan{
		token: token,
	}
}
