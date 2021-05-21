package una

var _ option = (*optionAkSk)(nil)

type optionAkSk struct {
	akSk akSk
}

// AkSk 配置创世云授权
func AkSk(ak string, sk string) *optionAkSk {
	return &optionAkSk{
		akSk: akSk{
			ak: ak,
			sk: sk,
		},
	}
}

func (as *optionAkSk) apply(options *options) {
	options.akSk = as.akSk
}
