package conf

import (
	"git.inke.cn/inkelogic/daenerys"
)

type SqlTblName struct {
	EsDealResTblName string `toml:"esDealRes"`
}

type Config struct {
	EsHost   EsHost     `toml:"esConfig`
	SqlTblName SqlTblName `toml:"sqlTblName`
}

type EsHost struct {
	PushHost     string `toml:"pushHost"`
	PullHost string `toml:"pullHost"`
	OpenHost string `toml:"openHost"`
}

var (
	Conf *Config
)

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := daenerys.ConfigInstance().Scan(cfg)
	Conf = cfg
	return cfg, err
}
