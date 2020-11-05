package conf

import (
	"git.inke.cn/inkelogic/daenerys"
)

type Config struct {
	EsConfig   EsConfig   `toml:"esConfig`
}

type EsConfig struct {
	ClsPushHost string `toml:"clsPushHost"`
	CdnPullOpenHost string `toml:"cdnPullOpenHost"`
}

var (
	Conf  *Config
)

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := daenerys.ConfigInstance().Scan(cfg)
	Conf = cfg
	return cfg, err
}
