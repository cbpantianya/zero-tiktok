package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	RPC struct {
		User struct {
			Host string `json:"host"`
			Port int64  `json:"port"`
		} `json:"user"`
	} `json:"rpc"`
}
