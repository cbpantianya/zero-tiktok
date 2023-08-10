package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-tiktok/api/internal/config"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"
)

type ServiceContext struct {
	Config      config.Config
	UserClient  user.UserServiceClient
	VideoClient video.VideoServiceClient
	Interaction interaction.InteractionServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	connU := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{fmt.Sprintf("%s:%d", c.RPC.User.Host, c.RPC.User.Port)},
	})

	connV := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{fmt.Sprintf("%s:%d", c.RPC.Video.Host, c.RPC.Video.Port)},
	})

	connI := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{fmt.Sprintf("%s:%d", c.RPC.Interaction.Host, c.RPC.Interaction.Port)},
	})

	uc := user.NewUserServiceClient(connU.Conn())

	vid := video.NewVideoServiceClient(connV.Conn())

	ia := interaction.NewInteractionServiceClient(connI.Conn())

	return &ServiceContext{
		Config:      c,
		UserClient:  uc,
		VideoClient: vid,
		Interaction: ia,
	}
}
