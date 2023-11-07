package qingniao

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/http"
	"github.com/goexl/log"

	"github.com/goexl/gox/field"
	"github.com/goexl/xiren"
)

var _ smsExecutor = (*chuangcache)(nil)

type chuangcache struct {
	ak          string
	sk          string
	apiEndpoint string
	smsEndpoint string
	token       *token

	code   chuangcacheCode
	http   *http.Client
	logger log.Logger
}

func newChuangcache(ak string, sk string, http *http.Client, logger log.Logger) *chuangcache {
	return &chuangcache{
		ak:          ak,
		sk:          sk,
		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",

		code:   newChuangcacheCode(),
		http:   http,
		logger: logger,
	}
}

func (c *chuangcache) send(ctx context.Context, deliver *smsDeliverInternal) (id string, status Status, err error) {
	if err = xiren.Struct(deliver); nil != err {
		return
	}

	baseReq := &baseChuangcacheSmsReq{
		AppKey:  deliver.Key.(string),
		Mobile:  strings.Join(deliver.Mobiles, ","),
		Content: deliver.Content,
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
	switch deliver.Type {
	case smsTypeCode:
		request.SetBody(chuangcacheOrdinaryReq{
			baseChuangcacheSmsReq: baseReq,
			Type:                  1,
		})
	case smsTypeNotify:
		request.SetBody(chuangcacheOrdinaryReq{
			baseChuangcacheSmsReq: baseReq,
			Type:                  2,
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

	fields := gox.Fields[any]{
		field.New("content", deliver.Content),
		field.New("mobiles", deliver.Mobiles),
		field.New("app.key", deliver.Key),
		field.New("id", id),
	}
	// 设置状态
	switch rsp.Code {
	case c.code.success:
		status = StatusAccepted
		c.logger.Debug("短信已提交到创世云", fields...)
	case c.code.badRequest:
		err = exc.NewException(rsp.Code, "请求参数错误", fields...)
	case c.code.unauthorized:
		err = exc.NewException(rsp.Code, "请求鉴权错误", fields...)
	case c.code.userNotfound:
		err = exc.NewException(rsp.Code, "用户不存在", fields...)
	case c.code.appNotfound:
		err = exc.NewException(rsp.Code, "短信服务不存在", fields...)
	case c.code.failed:
		status = StatusReject
		c.logger.Warn("短信提交到创世去出错", fields...)
	case c.code.noBalance:
		err = exc.NewException(rsp.Code, "余额不足", fields...)
	case c.code.timestampError:
		err = exc.NewException(rsp.Code, "时间戳错误", fields...)
	case c.code.tokenError:
		err = exc.NewException(rsp.Code, "授权码错误", fields...)
	case c.code.mobileInvalid:
		err = exc.NewException(rsp.Code, "存在不合法手机号", fields...)
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
	if hr, pe := c.http.R().SetContext(ctx).SetBody(req).SetResult(rsp).Post(url); nil != pe {
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
