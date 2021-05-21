package una

import (
	`context`
	`fmt`
	`strconv`
	`time`

	`github.com/go-resty/resty/v2`
	`github.com/mcuadros/go-defaults`
	`github.com/storezhang/gox`
	`github.com/storezhang/validatorx`
)

type (
	// ChuangcacheSmsConfig 创世云短信通知配置
	ChuangcacheSmsConfig struct {
		// 授权
		Secret gox.Secret `json:"secret" yaml:"secret" validate:"required"`
	}

	// ChuangcacheSms 创世云短信
	ChuangcacheSms struct {
		config ChuangcacheSmsConfig
		resty  *resty.Request

		token     string
		expiresIn time.Time

		apiEndpoint string
		smsEndpoint string
	}

	// 响应基类
	baseChuangcacheResponse struct {
		// 接口返回码，0：操作失败；1：操作成功
		Status int `json:"status"`
		// 接口返回信息，操作失败/操作成功
		Info string `json:"info"`
	}

	// 刷新Token请求
	chuangcacheTokenRequest struct {
		// 授权，类似于用户名
		Ak string `json:"ak"`
		// 授权，类似于密码
		Sk string `json:"sk"`
	}

	// 刷新Token响应
	chuangcacheTokenResponse struct {
		baseChuangcacheResponse

		//  接口返回数据对象
		Data struct {
			// 用户授权的唯一票据，用于调用接口的唯一票据
			AccessToken string `json:"access_token"`
			//  AccessToken的生命周期，单位是秒数
			ExpiresIn int `json:"expires_in"`
		} `json:"data"`
	}

	// 短信请求基类
	baseChuangcacheSmsRequest struct {
		// 密钥
		AccessToken string `json:"access_token"`
		// 短信服务标识
		AppKey string `json:"app_key"`
		// 合法的手机号码
		Mobile string `json:"mobile"`
		// 短信内容，长度不能超过536个字符，使用URL方式编码为UTF-8格式。短信内容超过70个字符时，会被拆分成多条，然后以长短信的格式发送
		Content string `json:"content"`
		// 时间戳（距离1970-1-1的毫秒数）
		Time string `json:"time"`
	}

	// 发送普通短信请求
	chuangcacheOrdinaryRequest struct {
		baseChuangcacheSmsRequest

		// 短信类型，默认为1，如果为通知短信则为2
		SmsType int `json:"sms_type"`
	}

	// 发送营销短信请求
	chuangcacheAdvertisingRequest struct {
		baseChuangcacheSmsRequest
	}

	// 发送短信响应
	chuangcacheSmsResponse struct {
		// 接口返回码
		Code int `json:"code"`
		// 发送短信流水号
		SendId string `json:"sendid"`
		// 接口返回信息
		Msg string `json:"msg"`
	}
)

// NewChuangcacheSms 创建创世云短信
func NewChuangcacheSms(config ChuangcacheSmsConfig, resty *resty.Request) (chuangcache *ChuangcacheSms, err error) {
	// 处理默认值
	defaults.SetDefaults(&config)
	if err = validatorx.Validate(config); nil != err {
		return
	}

	chuangcache = &ChuangcacheSms{
		config: config,
		resty:  resty,

		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "http://sms.chuangcache.com/api/sms",
	}

	return
}

func (cs *ChuangcacheSms) Send(_ context.Context, to string, content string, opts ...option) (id string, err error) {
	appliedOptions := defaultOptions()
	for _, opt := range opts {
		opt.apply(appliedOptions)
	}

	if err = cs.refreshToken(); nil != err {
		return
	}

	baseReq := baseChuangcacheSmsRequest{
		AccessToken: cs.token,
		AppKey:      cs.config.Secret.Id,
		Mobile:      to,
		Content:     content,
		Time:        strconv.FormatInt(time.Now().Unix(), 10),
	}

	var req interface{}
	switch appliedOptions.smsType {
	case SmsTypeCommon:
		fallthrough
	case SmsTypeNotify:
		req = chuangcacheOrdinaryRequest{
			baseChuangcacheSmsRequest: baseReq,
			SmsType:                   int(appliedOptions.smsType),
		}
	case SmsTypeAdvertising:
		req = chuangcacheAdvertisingRequest{
			baseChuangcacheSmsRequest: baseReq,
		}
	}

	rsp := new(chuangcacheSmsResponse)
	url := fmt.Sprintf("%s/%s", cs.smsEndpoint, "ordinary")
	if _, err = cs.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}
	id = rsp.SendId

	return
}

func (cs *ChuangcacheSms) refreshToken() (err error) {
	now := time.Now()
	if now.After(cs.expiresIn.Add(5 * time.Minute)) {
		return
	}

	// 更新Token
	req := chuangcacheTokenRequest{
		Ak: cs.config.Secret.Id,
		Sk: cs.config.Secret.Key,
	}
	rsp := new(chuangcacheTokenResponse)
	url := fmt.Sprintf("%s/%s", cs.apiEndpoint, "OAuth/authorize")
	if _, err = cs.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}
	cs.token = rsp.Data.AccessToken
	cs.expiresIn = time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn))

	return
}
