package qingniao_test

import (
	"context"
	"testing"

	"github.com/goexl/qingniao"
)

func TestValidate(test *testing.T) {
	testcases := []struct {
		template    string
		content     string
		mobiles     []string
		exceptError bool
	}{
		{"", "", []string{"17089792784"}, true},
		{"测试", "测试", []string{"17089792784"}, false},
	}

	ctx := context.Background()
	chuangcache := qingniao.New().Sms().Chuangcache("", "")
	for index, testcase := range testcases {
		_, _, err := chuangcache.Deliver(testcase.template, testcase.content, testcase.mobiles...).Send(ctx)
		switch {
		case nil == err && testcase.exceptError:
			test.Errorf("第%d个测试出错，应该报数据验证不通过", index+1)
		case nil != err && !testcase.exceptError:
			test.Errorf("第%d个测试出错，数据验证应该通过", index+1)
		}
	}
}
