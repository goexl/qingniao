package executor

import (
	"context"
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	internal2 "github.com/goexl/qingniao/internal/kernel"
)

var _ internal.Wechat = (*ServerChan)(nil)

type ServerChan struct {
	key string

	http   *http.Client
	logger log.Logger
}

func NewServerChan(key string, http *http.Client, logger log.Logger) *ServerChan {
	return &ServerChan{
		key: key,

		http:   http,
		logger: logger,
	}
}

func (sc *ServerChan) Send(ctx context.Context, deliver *deliver.Wechat) (id string, status internal2.Status, err error) {
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
		status = internal2.StatusDelivered
	}

	return
}
