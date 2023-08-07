package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-tiktok/api/internal/config"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
)

type ServiceContext struct {
	Config     config.Config
	UserClient user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{fmt.Sprintf("%s:%d", c.RPC.User.Host, c.RPC.User.Port)},
	})

	uc := user.NewUserServiceClient(conn.Conn())
	return &ServiceContext{
		Config:     c,
		UserClient: uc,
	}
}
