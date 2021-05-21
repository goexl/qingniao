package una

const (
	// EmailTypeHtml 富文本邮件
	EmailTypeHtml EmailType = "html"
	// EmailTypePlain 普通文本邮件
	EmailTypePlain EmailType = "plain"
)

// EmailType 邮件类型
type EmailType string
