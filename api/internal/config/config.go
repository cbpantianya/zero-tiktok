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
	OSS struct {
		Endpoint        string `json:"endpoint"`
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		BucketName      string `json:"bucket"`
		Domain          string `json:"domain"`
	} `json:"oss"`
}
