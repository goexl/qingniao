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

var _ internal.Wechat = (*Serverchan)(nil)

type Serverchan struct {
	key string

	http   *http.Client
	logger log.Logger
}

func NewServerchan(key string, http *http.Client, logger log.Logger) *Serverchan {
	return &Serverchan{
		key: key,

		http:   http,
		logger: logger,
	}
}

func (s *Serverchan) Send(ctx context.Context, deliver *deliver.Wechat) (id string, status internal2.Status, err error) {
	form := map[string]string{
		"title": deliver.Title,
		"desp":  deliver.Content,
	}
	url := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", s.key)
	if hr, pe := s.http.R().SetContext(ctx).SetFormData(form).Post(url); nil != pe {
		err = pe
	} else if hr.IsError() {
		s.logger.Warn("Serverchan返回错误", field.New("status.code", hr.StatusCode()))
	} else {
		status = internal2.StatusDelivered
	}

	return
}
