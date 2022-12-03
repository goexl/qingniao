package qingniao

type (
	// 刷新请求
	chuangcacheTokenReq struct {
		// 授权，类似于用户名
		Ak string `json:"ak"`
		// 授权，类似于密码
		Sk string `json:"sk"`
	}

	// 刷新响应
	chuangcacheTokenRsp struct {
		baseChuangcacheRsp

		//  接口返回数据对象
		Data struct {
			// 用户授权的唯一票据，用于调用接口的唯一票据
			AccessToken string `json:"access_token"`
			//  AccessToken的生命周期，单位是秒数
			ExpiresIn int `json:"expires_in,string"`
		} `json:"data"`
	}
)
