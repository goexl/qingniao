package una

import (
	`fmt`
)

type akSk struct {
	// 授权，类似于用户名
	ak string `validate:"required"`
	// 授权，类似于密码
	sk string `validate:"required"`
}

func (as *akSk) key() string {
	return fmt.Sprintf("%s-%s", as.ak, as.sk)
}
