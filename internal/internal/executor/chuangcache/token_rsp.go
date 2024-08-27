package chuangcache

type TokenRsp struct {
	baseRsp

	// 接口返回数据对象
	Data struct {
		// 用户授权的唯一票据，用于调用接口的唯一票据
		AccessToken string `json:"access_token"`
		//  AccessToken的生命周期，单位是秒数
		ExpiresIn int `json:"expires_in,string"`
	} `json:"data"`
}
