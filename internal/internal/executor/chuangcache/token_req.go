package chuangcache

type TokenReq struct {
	// 授权，类似于用户名
	Ak string `json:"ak"`
	// 授权，类似于密码
	Sk string `json:"sk"`
}
