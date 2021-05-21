package una

// Config 配置
type Config struct {
	// 类型
	Type Type `json:"type" yaml:"type" validate:"required,oneof=cos"`
	// 创世云短信
	ChuangcacheSms ChuangcacheSmsConfig `json:"chuangcacheSms" yaml:"chuangcacheSms" validate:"structonly"`
}
