package logic

import (
	"context"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个用户信息
func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserResponse{}, nil
}
