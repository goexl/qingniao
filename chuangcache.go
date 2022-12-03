package qingniao

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/goexl/xiren"
)

var _ smsExecutor = (*chuangcache)(nil)

type chuangcache struct {
	ak          string
	sk          string
	apiEndpoint string
	smsEndpoint string
	token       *token

	http   *resty.Client
	logger simaqian.Logger
}

func newChuangcache(ak string, sk string, http *resty.Client, logger simaqian.Logger) *chuangcache {
	return &chuangcache{
		ak:          ak,
		sk:          sk,
		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",

		http:   http,
		logger: logger,
	}
}

func (c *chuangcache) send(ctx context.Context, deliver *smsDeliver) (id string, err error) {
	if err = xiren.Struct(deliver); nil != err {
		return
	}

	baseReq := &baseChuangcacheSmsReq{
		AppKey:  "",
		Mobile:  strings.Join(deliver.mobiles, ","),
		Content: deliver.content,
		Time:    fmt.Sprintf("%d", time.Now().UnixNano()/1e6),
	}
	if _token, te := c.getToken(ctx); nil != te {
		err = te
	} else {
		baseReq.AccessToken = _token
	}
	if nil != err {
		return
	}

	request := c.http.R()
	switch deliver.typ {
	case smsTypeCommon:
		fallthrough
	case smsTypeNotify:
		request.SetBody(chuangcacheOrdinaryReq{
			baseChuangcacheSmsReq: baseReq,
			Type:                  int(deliver.typ),
		})
	case smsTypeAdvertising:
		request.SetBody(chuangcacheAdvertisingReq{
			baseChuangcacheSmsReq: baseReq,
		})
	}

	rsp := new(chuangcacheSmsRsp)
	url := fmt.Sprintf("%s/%s", c.smsEndpoint, "ordinary")
	if hr, pe := request.SetContext(ctx).SetResult(rsp).Post(url); nil != pe {
		err = pe
	} else if hr.IsError() {
		c.logger.Warn("创世云返回错误", field.New("status.code", hr.StatusCode()))
	} else {
		id = rsp.Id
	}

	return
}

func (c *chuangcache) getToken(ctx context.Context) (_token string, err error) {
	if nil != c.token && c.token.validate() {
		_token = c.token.token
	}
	if "" != _token {
		return
	}

	req := chuangcacheTokenReq{
		Ak: c.ak,
		Sk: c.sk,
	}
	rsp := new(chuangcacheTokenRsp)
	url := fmt.Sprintf("%s/%s", c.apiEndpoint, "OAuth/authorize")
	if hr, pe := c.http.R().SetContext(ctx).SetBody(req).SetResult(rsp).Post(url); nil == pe {
		err = pe
	} else if hr.IsError() {
		c.logger.Warn("创世云返回错误", field.New("status.code", hr.StatusCode()))
	} else {
		c.token = new(token)
		c.token.token = rsp.Data.AccessToken
		c.token.expiresIn = time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn))
		_token = rsp.Data.AccessToken
	}

	return
}
