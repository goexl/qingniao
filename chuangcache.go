package una

import (
	`context`
	`fmt`
	`strings`
	`sync`
	`time`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/validatorx`
)

// Chuangcache 创世云短信
type Chuangcache struct {
	resty      *resty.Request
	tokenCache sync.Map

	apiEndpoint string
	smsEndpoint string

	template unaTemplate
}

// NewChuangcache 创建创世云短信
func NewChuangcache(resty *resty.Request) (chuangcache *Chuangcache) {
	chuangcache = &Chuangcache{
		resty:      resty,
		tokenCache: sync.Map{},

		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",
	}
	chuangcache.template = unaTemplate{chuangcache: chuangcache}

	return
}

func (c *Chuangcache) Send(ctx context.Context, content string, opts ...option) (id string, err error) {
	return c.template.Send(ctx, content, opts...)
}

func (c *Chuangcache) send(_ context.Context, content string, options *options) (id string, err error) {
	if err = validatorx.Var(content, "required,max=536"); nil != err {
		return
	}
	if err = validatorx.Struct(options.chuangcache); nil != err {
		return
	}

	var token string
	if token, err = c.refreshToken(options); nil != err {
		return
	}

	baseReq := baseChuangcacheSmsRequest{
		AccessToken: token,
		AppKey:      options.chuangcache.appKey,
		Mobile:      strings.Join(options.chuangcache.mobiles, ","),
		Content:     content,
		Time:        fmt.Sprintf("%d", time.Now().UnixNano()/1e6),
	}

	var req interface{}
	switch options.chuangcache.smsType {
	case SmsTypeCommon:
		fallthrough
	case SmsTypeNotify:
		req = chuangcacheOrdinaryRequest{
			baseChuangcacheSmsRequest: baseReq,
			SmsType:                   int(options.chuangcache.smsType),
		}
	case SmsTypeAdvertising:
		req = chuangcacheAdvertisingRequest{
			baseChuangcacheSmsRequest: baseReq,
		}
	}

	rsp := new(chuangcacheSmsResponse)
	url := fmt.Sprintf("%s/%s", c.smsEndpoint, "ordinary")
	if _, err = c.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}
	id = rsp.SendId

	return
}

func (c *Chuangcache) refreshToken(options *options) (token string, err error) {
	var (
		cache interface{}
		ok    bool
	)

	key := options.chuangcache.key()
	// 检查AccessToken是否可以
	if cache, ok = c.tokenCache.Load(key); ok {
		var validate bool
		if token, validate = cache.(*chuangcacheToken).validate(); validate {
			return
		}
	}

	// 更新Token
	req := chuangcacheTokenRequest{
		Ak: options.chuangcache.ak,
		Sk: options.chuangcache.sk,
	}
	rsp := new(chuangcacheTokenResponse)
	url := fmt.Sprintf("%s/%s", c.apiEndpoint, "OAuth/authorize")
	if _, err = c.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}

	token = rsp.Data.AccessToken
	c.tokenCache.Store(key, &chuangcacheToken{
		token:     token,
		expiresIn: time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn)),
	})

	return
}
