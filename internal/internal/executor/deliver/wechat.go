package deliver

type Wechat struct {
	Title   string `validate:"required"`
	Content string `validate:"required"`
}
