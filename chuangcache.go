package una

import (
	`context`
	`fmt`
	`strconv`
	`strings`
	`sync`
	`time`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/validatorx`
)

// ChuangcacheSms 创世云短信
type ChuangcacheSms struct {
	resty      *resty.Request
	tokenCache sync.Map

	apiEndpoint string
	smsEndpoint string
}

// NewChuangcacheSms 创建创世云短信
func NewChuangcacheSms(resty *resty.Request) (chuangcache *ChuangcacheSms, err error) {
	chuangcache = &ChuangcacheSms{
		resty:      resty,
		tokenCache: sync.Map{},

		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",
	}

	return
}

func (cs *ChuangcacheSms) Send(_ context.Context, content string, opts ...option) (id string, err error) {
	appliedOptions := defaultOptions()
	for _, opt := range opts {
		opt.apply(appliedOptions)
	}

	var token string
	if token, err = cs.refreshToken(appliedOptions); nil != err {
		return
	}

	baseReq := baseChuangcacheSmsRequest{
		AccessToken: token,
		AppKey:      appliedOptions.akSk.ak,
		Mobile:      strings.Join(appliedOptions.sms.targets, ","),
		Content:     content,
		Time:        strconv.FormatInt(time.Now().Unix(), 10),
	}

	var req interface{}
	switch appliedOptions.sms.Type {
	case SmsTypeCommon:
		fallthrough
	case SmsTypeNotify:
		req = chuangcacheOrdinaryRequest{
			baseChuangcacheSmsRequest: baseReq,
			SmsType:                   int(appliedOptions.sms.Type),
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

func (cs *ChuangcacheSms) refreshToken(options *options) (token string, err error) {
	if err = validatorx.Validate(options.akSk); nil != err {
		return
	}

	var (
		cache interface{}
		ok    bool
	)
	// 检查AccessToken是否可以
	if cache, ok = cs.tokenCache.Load(options.akSk.key()); ok {
		var validate bool
		if token, validate = cache.(*chuangcacheToken).validate(); validate {
			return
		}
	}

	// 更新Token
	req := chuangcacheTokenRequest{
		Ak: options.akSk.ak,
		Sk: options.akSk.sk,
	}
	rsp := new(chuangcacheTokenResponse)
	url := fmt.Sprintf("%s/%s", cs.apiEndpoint, "OAuth/authorize")
	if _, err = cs.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}

	token = rsp.Data.AccessToken
	cs.tokenCache.Store(options.akSk.key(), &chuangcacheToken{
		token:     token,
		expiresIn: time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn)),
	})

	return
}
