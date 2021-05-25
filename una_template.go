package una

import (
	`bytes`
	`context`
	`text/template`
)

type unaTemplate struct {
	email       unaInternal
	chuangcache unaInternal
}

func (t *unaTemplate) Send(ctx context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	// 处理模板
	if options.template {
		var tpl *template.Template
		if tpl, err = template.New("una").Parse(content); nil != err {
			return
		}

		var buffer bytes.Buffer
		if err = tpl.Execute(&buffer, options.data); err != nil {
			return
		}

		content = buffer.String()
	}

	switch options.unaType {
	case TypeEmail:
		id, err = t.email.send(ctx, content, options)
	case TypeChuangcache:
		id, err = t.chuangcache.send(ctx, content, options)
	}

	return
}
