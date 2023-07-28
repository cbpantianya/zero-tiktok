package logic

import (
	"context"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取多个用户信息
func (l *GetUsersLogic) GetUsers(in *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUsersResponse{}, nil
}
