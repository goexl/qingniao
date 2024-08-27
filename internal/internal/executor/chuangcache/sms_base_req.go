package chuangcache

type SmsBaseReq struct {
	// 密钥
	AccessToken string `json:"access_token"`
	// 短信服务标识
	AppKey string `json:"app_key"`
	// 合法的手机号码
	Mobile string `json:"mobile"`
	// 短信内容
	// 长度不能超过536个字符，使用URL方式编码为UTF-8
	// 短信内容超过70个字符时，会被拆分成多条后以长短信的格式发送
	Content string `json:"content"`
	// 时间戳
	// 距离1970-1-1的毫秒数
	Time string `json:"time"`
}
