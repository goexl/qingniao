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

// Chuangcache 创世云短信
type Chuangcache struct {
	validate   *validatorx.Validate
	resty      *resty.Request
	tokenCache sync.Map

	apiEndpoint string
	smsEndpoint string
}

// NewChuangcache 创建创世云短信
func NewChuangcache(validate *validatorx.Validate, resty *resty.Request) *Chuangcache {
	return &Chuangcache{
		validate:   validate,
		resty:      resty,
		tokenCache: sync.Map{},

		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",
	}
}

func (cs *Chuangcache) Send(_ context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	if err = cs.validate.Var(content, "required,max=536"); nil != err {
		return
	}
	if err = cs.validate.Struct(options.chuangcache); nil != err {
		return
	}

	var token string
	if token, err = cs.refreshToken(options); nil != err {
		return
	}

	baseReq := baseChuangcacheSmsRequest{
		AccessToken: token,
		AppKey:      options.chuangcache.ak,
		Mobile:      strings.Join(options.chuangcache.mobiles, ","),
		Content:     content,
		Time:        strconv.FormatInt(time.Now().Unix(), 10),
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
	url := fmt.Sprintf("%s/%s", cs.smsEndpoint, "ordinary")
	if _, err = cs.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}
	id = rsp.SendId

	return
}

func (cs *Chuangcache) refreshToken(options *options) (token string, err error) {
	var (
		cache interface{}
		ok    bool
	)

	key := options.chuangcache.key()
	// 检查AccessToken是否可以
	if cache, ok = cs.tokenCache.Load(key); ok {
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
	url := fmt.Sprintf("%s/%s", cs.apiEndpoint, "OAuth/authorize")
	if _, err = cs.resty.SetBody(req).SetResult(rsp).Post(url); nil != err {
		return
	}

	token = rsp.Data.AccessToken
	cs.tokenCache.Store(key, &chuangcacheToken{
		token:     token,
		expiresIn: time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn)),
	})

	return
}
