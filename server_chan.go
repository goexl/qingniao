package qingniao

import (
	"context"
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/log"
)

var _ wechatExecutor = (*serverChan)(nil)

type serverChan struct {
	key string

	http   *http.Client
	logger log.Logger
}

func newServerChan(key string, http *http.Client, logger log.Logger) *serverChan {
	return &serverChan{
		key: key,

		http:   http,
		logger: logger,
	}
}

func (sc *serverChan) send(ctx context.Context, deliver *wechatDeliver) (id string, status Status, err error) {
	form := map[string]string{
		"title": deliver.Title,
		"desp":  deliver.Content,
	}
	url := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", sc.key)
	if hr, pe := sc.http.R().SetContext(ctx).SetFormData(form).Post(url); nil != pe {
		err = pe
	} else if hr.IsError() {
		sc.logger.Warn("ServerChan返回错误", field.New("status.code", hr.StatusCode()))
	} else {
		status = StatusDelivered
	}

	return
}
