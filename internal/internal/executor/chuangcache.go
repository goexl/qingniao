package executor

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/core"
	"github.com/goexl/qingniao/internal/internal/executor/chuangcache"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"

	"github.com/goexl/gox/field"
	"github.com/goexl/xiren"
)

var _ internal.Sms = (*Chuangcache)(nil)

type Chuangcache struct {
	ak          string
	sk          string
	apiEndpoint string
	smsEndpoint string
	token       *core.Token

	code   chuangcache.Code
	http   *http.Client
	logger log.Logger
}

func NewChuangcache(ak string, sk string, http *http.Client, logger log.Logger) *Chuangcache {
	return &Chuangcache{
		ak:          ak,
		sk:          sk,
		apiEndpoint: "https://api.chuangcache.com",
		smsEndpoint: "https://sms.chuangcache.com/api/sms",

		code:   chuangcache.NewCode(),
		http:   http,
		logger: logger,
	}
}

func (c *Chuangcache) Send(ctx context.Context, deliver *deliver.Sms) (id string, status kernel.Status, err error) {
	if se := xiren.Struct(deliver); nil != se {
		err = se
	} else if ce := c.check(deliver.Mobiles); nil != ce {
		err = ce
	}
	if nil != err {
		return
	}

	baseReq := &chuangcache.SmsBaseReq{
		AppKey:  deliver.Key,
		Mobile:  strings.ReplaceAll(strings.Join(deliver.Mobiles, ","), "+86", ""),
		Content: deliver.Content,
		Time:    fmt.Sprintf("%d", time.Now().UnixNano()/1e6),
	}
	if token, te := c.getToken(ctx); nil != te {
		err = te
	} else {
		baseReq.AccessToken = token
	}
	if nil != err {
		return
	}

	request := c.http.R()
	switch deliver.Type {
	case constant.SmsTypeCode:
		request.SetBody(chuangcache.SmsOrdinaryReq{
			SmsBaseReq: baseReq,
			Type:       1,
		})
	case constant.SmsTypeNotify:
		request.SetBody(chuangcache.SmsOrdinaryReq{
			SmsBaseReq: baseReq,
			Type:       2,
		})
	case constant.SmsTypeAdvertising:
		request.SetBody(chuangcache.SmsAdvertisingReq{
			SmsBaseReq: baseReq,
		})
	}

	rsp := new(chuangcache.SmsRsp)
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
	case c.code.Success:
		status = kernel.StatusAccepted
		c.logger.Debug("短信已提交到创世云", fields...)
	case c.code.BadRequest:
		err = exc.NewException(rsp.Code, "请求参数错误", fields...)
	case c.code.Unauthorized:
		err = exc.NewException(rsp.Code, "请求鉴权错误", fields...)
	case c.code.UserNotfound:
		err = exc.NewException(rsp.Code, "用户不存在", fields...)
	case c.code.AppNotfound:
		err = exc.NewException(rsp.Code, "短信服务不存在", fields...)
	case c.code.Failed:
		status = kernel.StatusReject
		c.logger.Warn("短信提交到创世去出错", fields...)
	case c.code.NoBalance:
		err = exc.NewException(rsp.Code, "余额不足", fields...)
	case c.code.TimestampError:
		err = exc.NewException(rsp.Code, "时间戳错误", fields...)
	case c.code.TokenError:
		err = exc.NewException(rsp.Code, "授权码错误", fields...)
	case c.code.MobileInvalid:
		err = exc.NewException(rsp.Code, "存在不合法手机号", fields...)
	}

	return
}

func (c *Chuangcache) getToken(ctx context.Context) (token string, err error) {
	if nil != c.token && c.token.Validate() {
		token = c.token.Code
	}
	if "" != token {
		return
	}

	req := chuangcache.TokenReq{
		Ak: c.ak,
		Sk: c.sk,
	}
	rsp := new(chuangcache.TokenRsp)
	url := fmt.Sprintf("%s/%s", c.apiEndpoint, "OAuth/authorize")
	if hr, pe := c.http.R().SetContext(ctx).SetBody(req).SetResult(rsp).Post(url); nil != pe {
		err = pe
	} else if hr.IsError() {
		c.logger.Warn("创世云返回错误", field.New("status.code", hr.StatusCode()))
	} else {
		c.token = new(core.Token)
		c.token.Code = rsp.Data.AccessToken
		c.token.ExpiresIn = time.Now().Add(time.Duration(1000 * rsp.Data.ExpiresIn))
		token = rsp.Data.AccessToken
	}

	return
}

func (c *Chuangcache) check(mobiles []string) (err error) {
	banlist := make([]string, 0)
	for _, mobile := range mobiles {
		if !strings.HasPrefix(mobile, constant.MobileChinaPrefix) {
			banlist = append(banlist, mobile)
		}
	}
	if 0 != len(banlist) {
		err = exception.New().Message("不支持除中国以外的手机号发送短信").Field(field.New("mobiles", banlist)).Build()
	}

	return
}
