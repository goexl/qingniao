package qingniao

import (
	"github.com/goexl/qingniao/internal/kernel"
)

const (
	// StatusUnknown 未知状态
	StatusUnknown = kernel.StatusUnknown

	// StatusAccepted 已接受
	StatusAccepted = kernel.StatusAccepted
	// StatusDelivering 投递中
	StatusDelivering = kernel.StatusDelivering

	// StatusDelivered 短消息发送成功
	StatusDelivered = kernel.StatusDelivered

	// StatusReject 审核驳回
	StatusReject = kernel.StatusReject
	// StatusExpired 短消息超过有效期
	StatusExpired = kernel.StatusExpired
	// StatusUnreachable 短消息是不可达的
	StatusUnreachable = kernel.StatusUnreachable
	// StatusRejected 短消息被中心拒绝
	StatusRejected = kernel.StatusRejected
	// StatusBlacklist 目标是黑名单号码
	StatusBlacklist = kernel.StatusBlacklist
	// StatusError 系统忙
	StatusError = kernel.StatusError
)

// Status 状态
type Status = kernel.Status
