package config

type Config struct {
	ServerPort int
}

var _cfg *Config

func Get() *Config {
	return _cfg
}

func Set(cfg *Config) {
	_cfg = cfg
}
