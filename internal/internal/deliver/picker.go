package deliver

import (
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/qingniao/internal/internal/constant"
)

type picker[T any] struct {
	executors map[constant.Executor]T
	first     T
	set       bool
}

func newPicker[T any](executors map[constant.Executor]T) *picker[T] {
	return &picker[T]{
		executors: executors,
	}
}

func (p *picker[T]) pick(current constant.Executor, name string) (executor T, err error) {
	if 0 != len(p.executors) {
		err = exception.New().Message("没有配置任何发送执行器").
			Field(field.New("executors", p.executors)).Field(field.New("type", name)).
			Build()
	} else if constant.ExecutorUnknown == current {
		executor = p.getFirst()
	} else if cached, ok := p.executors[current]; !ok {
		err = exception.New().Message("没有配置发送执行器").
			Field(field.New("executors", p.executors), field.New("current", current)).Field(field.New("type", name)).
			Build()
	} else {
		executor = cached
	}

	return
}

func (p *picker[T]) getFirst() T {
	if !p.set {
		for _, executor := range p.executors {
			p.first = executor
			p.set = true

			if p.set {
				break
			}
		}
	}

	return p.first
}
