package qingniao

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/storezhang/validatorx"
)

var _ smsExecutor = (*chuangcache)(nil)

type chuangcache struct {
	http       *resty.Client
	tokenCache sync.Map

	ak          string
	sk          string
	apiEndpoint string
	smsEndpoint string
	token       *chuangcacheToken
}

func newChuangcache(ak string, sk string, http *resty.Client) *chuangcache {
	return &chuangcache{
		http:       http,
		tokenCache: sync.Map{},

		ak:          ak,
		sk:          sk,
		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",
	}
}

func (c *chuangcache) Send(ctx context.Context) (id string, err error) {
	if err = validatorx.Var(content, "required,max=536"); nil != err {
		return
	}
	if err = validatorx.Struct(options.chuangcache); nil != err {
		return
	}

	var token string
	baseReq := new(baseChuangcacheSmsRequest)
	if token, err = c.getToken(ctx); nil != err {
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
	if _, err = c.http.R().SetContext(ctx).SetBody(req).SetResult(rsp).Post(url); nil == err {
		id = rsp.SendId
	}

	return
}

func (c *chuangcache) getToken(ctx context.Context) (token string, err error) {
	if nil != c.token && c.token.validate() {
		token = c.token.token
	}
	if "" != token {
		return
	}

	req := chuangcacheTokenRequest{
		Ak: c.ak,
		Sk: c.sk,
	}
	rsp := new(chuangcacheTokenResponse)
	url := fmt.Sprintf("%s/%s", c.apiEndpoint, "OAuth/authorize")
	if _, err = c.http.R().SetContext(ctx).SetBody(req).SetResult(rsp).Post(url); nil == err {
		c.token = new(chuangcacheToken)
		c.token.token = rsp.Data.AccessToken
		c.token.expiresIn = time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn))
		token = rsp.Data.AccessToken
	}

	return
}
