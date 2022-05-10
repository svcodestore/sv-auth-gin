package config

type System struct {
	Id            string `mapstructure:"id" json:"id" yaml:"id"`
	Secret        string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	SsoRpcAddr    string `mapstructure:"sso-rpc-addr" json:"ssoRpcAddr" yaml:"sso-rpc-addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
}
