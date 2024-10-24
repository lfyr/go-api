package config

type System struct {
	Env         string `mapstructure:"env" json:"env" yaml:"env"`                            // 环境值
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`                         // 端口值
	TokenExpire int    `mapstructure:"token_expire" json:"token_expire" yaml:"token_expire"` // 端口值
}

type Server struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
