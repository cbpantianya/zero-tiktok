package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	RPC struct {
		User struct {
			Host string `json:"host"`
			Port int64  `json:"port"`
		} `json:"user"`
		Video struct {
			Host string `json:"host"`
			Port int64  `json:"port"`
		} `json:"video"`

		Interaction struct {
			Host string `json:"host"`
			Port int64  `json:"port"`
		} `json:"interaction"`
	} `json:"rpc"`
}
