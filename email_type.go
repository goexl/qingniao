package qingniao

const (
	// emailTypeHtml 富文本邮件
	emailTypeHtml emailType = "html"
	// emailTypePlain 普通文本邮件
	emailTypePlain emailType = "plain"
)

// emailType 邮件类型
type emailType string
