package qingniao

// 响应基类
type baseChuangcacheRsp struct {
	// 接口返回码，0：操作失败；1：操作成功
	Status int `json:"status"`
	// 接口返回信息，操作失败/操作成功
	Info string `json:"info"`
}
