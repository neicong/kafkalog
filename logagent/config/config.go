package config

import (
	"github.com/go-ini/ini"
)

type KafkaConfig struct {
	Addrress string `ini:"addrress"`
	Topic string `ini:"topic"`
	MsgChanSize int64 `ini:"msg_chan_size"`
}

type CollectConfig struct {
	LogfilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	EtcdAddrress string `ini:"etcd_addrress"`
}

type Config struct {
	KafkaConfig `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig `ini:"etcd"`
}

func GetConfig() *Config{
	var configObj = new(Config)
	ini.MapTo(configObj,"./config/config.ini")
	return configObj
}

