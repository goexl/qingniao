package una

var _ option = (*optionAkSk)(nil)

type optionAkSk struct {
	// 授权，相当于用户名
	ak string
	// 授权，相当于密码
	sk string
}

// AkSk 配置邮件服务
func AkSk(ak string, sk string) *optionAkSk {
	return &optionAkSk{
		ak: ak,
		sk: sk,
	}
}

func (as *optionAkSk) apply(options *options) {
	options.chuangcacheSms.ak = as.ak
	options.chuangcacheSms.sk = as.sk
}
