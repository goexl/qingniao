package una

import (
	`fmt`
)

type chuangcacheConfig struct {
	// 授权，相当于用户名
	ak string `validate:"required"`
	// 授权，相当于密码
	sk string `validate:"required"`
	// 应用
	appKey string `validate:"required"`
	// 短信类型
	smsType SmsType `validate:"required,oneof=1 2 3"`
	// 手机号列表
	mobiles []string `validate:"required,dive,mobile"`
}

func (cc *chuangcacheConfig) key() string {
	return fmt.Sprintf("%s-%s", cc.ak, cc.sk)
}
