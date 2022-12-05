package qingniao

var (
	_ = StatusUnknown
	_ = StatusAccepted
	_ = StatusDelivering
	_ = StatusDelivered
	_ = StatusReject
	_ = StatusExpired
	_ = StatusUnreachable
	_ = StatusRejected
	_ = StatusBlacklist
	_ = StatusError
)

const (
	// StatusUnknown 未知状态
	StatusUnknown Status = 0

	// StatusAccepted 已接受
	StatusAccepted Status = 1
	// StatusDelivering 投递中
	StatusDelivering Status = 2

	// StatusDelivered 短消息发送成功
	StatusDelivered Status = 10

	// StatusReject 审核驳回
	StatusReject Status = 20
	// StatusExpired 短消息超过有效期
	StatusExpired Status = 21
	// StatusUnreachable 短消息是不可达的
	StatusUnreachable Status = 22
	// StatusRejected 短消息被中心拒绝
	StatusRejected Status = 23
	// StatusBlacklist 目标是黑名单号码
	StatusBlacklist Status = 24
	// StatusError 系统忙
	StatusError Status = 25
)

// Status 状态
type Status uint8
