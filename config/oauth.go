package config

type Oauth struct {
	TokenUrl string `mapstructure:"token-url" json:"token-url" yaml:"token-url"`
}
